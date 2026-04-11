# ELF Binary Extractor

> **Disclaimer**
> This project is a fork of an existing implementation. I am not the original author.
> I have only refactored and cleaned the code for readability and structure.
>
> Full credit belongs to the original creator.


```py
import sys, struct, os
from elftools.elf.elffile import ELFFile

def vaddr_to_offset(elf, vaddr):
    for segment in elf.iter_segments():
        if segment['p_type'] == 'PT_LOAD':
            start = segment['p_vaddr']
            end = start + segment['p_memsz']
            if start <= vaddr < end:
                return (vaddr - start) + segment['p_offset']
    raise ValueError(hex(vaddr))

def read_go_string(stream, elf, ptr, length):
    if length == 0:
        return b""
    try:
        offset = vaddr_to_offset(elf, ptr)
        stream.seek(offset)
        return stream.read(length)
    except ValueError:
        return b"<invalid_ptr>"

def find_embed_symbol(elf):
    symbol_table = elf.get_section_by_name('.symtab')
    if not symbol_table:
        print("No symbol table found.")
        return None

    keywords = ['embed', 'assets', 'static', 'content', 'fs', 'files']
    candidates = []

    for sym in symbol_table.iter_symbols():
        if sym['st_info']['type'] == 'STT_OBJECT' and sym['st_size'] == 8:
            name = sym.name.lower()
            if any(x in name for x in ('runtime.', 'type.', 'go.')):
                continue
            if any(k in name for k in keywords):
                candidates.append(sym)

    if not candidates:
        return None

    best = candidates[0]
    print(f"Auto-selected: {best.name} @ {hex(best['st_value'])}")
    return best['st_value']

def main():
    if len(sys.argv) < 2:
        print("Usage: python3 extract_embed.py <binary_path> [hex_addr]")
        sys.exit(1)

    path = sys.argv[1]
    with open(path, 'rb') as f:
        elf = ELFFile(f)

        if len(sys.argv) >= 3:
            fs_vaddr = int(sys.argv[2], 16)
        else:
            print("Attempting to auto-detect embed.FS...")
            fs_vaddr = find_embed_symbol(elf)
            if fs_vaddr is None:
                print("Auto-detect failed. Provide address manually.")
                sys.exit(1)

        try:
            fs_offset = vaddr_to_offset(elf, fs_vaddr)
        except ValueError as e:
            print(f"Address error: {e}")
            sys.exit(1)

        f.seek(fs_offset)
        slice_header_vaddr = struct.unpack('<Q', f.read(8))[0]
        if slice_header_vaddr == 0:
            print("Pointer is NULL.")
            sys.exit(1)

        try:
            slice_header_offset = vaddr_to_offset(elf, slice_header_vaddr)
        except ValueError:
            print("Slice header points to invalid memory.")
            sys.exit(1)

        f.seek(slice_header_offset)
        files_array_vaddr = struct.unpack('<Q', f.read(8))[0]
        files_count = struct.unpack('<Q', f.read(8))[0]

        print(f"Files array at {hex(files_array_vaddr)}, count {files_count}")

        if files_count > 10000:
            print("Suspiciously high count. Abort.")
            sys.exit(1)

        FILE_STRUCT_SIZE = 48
        output_dir = "extracted_embed"
        os.makedirs(output_dir, exist_ok=True)

        current_file_vaddr = files_array_vaddr

        for i in range(files_count):
            try:
                file_offset = vaddr_to_offset(elf, current_file_vaddr)
            except ValueError:
                print(f"Error reading file struct #{i}. Stopping.")
                break

            f.seek(file_offset)
            name_ptr = struct.unpack('<Q', f.read(8))[0]
            name_len = struct.unpack('<Q', f.read(8))[0]
            data_ptr = struct.unpack('<Q', f.read(8))[0]
            data_len = struct.unpack('<Q', f.read(8))[0]

            file_name_bytes = read_go_string(f, elf, name_ptr, name_len)
            file_data = read_go_string(f, elf, data_ptr, data_len)
            file_name = file_name_bytes.decode('utf-8', errors='ignore') or f"unknown_{i}"

            clean_name = file_name.replace('..', '')
            if clean_name.startswith('/'):
                clean_name = clean_name[1:]

            full_path = os.path.join(output_dir, clean_name)
            os.makedirs(os.path.dirname(full_path), exist_ok=True)

            print(f"Extracting: {file_name} ({data_len} bytes)")

            if not os.path.isdir(full_path) and data_len > 0:
                with open(full_path, 'wb') as out_f:
                    out_f.write(file_data)

            current_file_vaddr += FILE_STRUCT_SIZE

        print(f"Extraction complete. Folder: {output_dir}")

if __name__ == "__main__":
    main()
```