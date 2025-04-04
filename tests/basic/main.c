int dereferenceOffsetInt(int* p, int offset) {
    return *(p + offset);
}

int main() {
    int a[5] = {1, 2, 3, 4, 5};
    int* p = a;
    int offset = 8;
    int result = dereferenceOffsetInt(p, offset);
    
    return result;
}