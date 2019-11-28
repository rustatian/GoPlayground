
#line 1 "cgo-c-prolog-gccgo"
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

typedef unsigned char byte;
typedef intptr_t intgo;

struct __go_string {
	const unsigned char *__data;
	intgo __length;
};

typedef struct __go_open_array {
	void* __values;
	intgo __count;
	intgo __capacity;
} Slice;

struct __go_string __go_byte_array_to_string(const void* p, intgo len);
struct __go_open_array __go_string_to_byte_array (struct __go_string str);

const char *_cgo_ca7c2c1b0b09_Cfunc_CString(struct __go_string s) {
	char *p = malloc(s.__length+1);
	memmove(p, s.__data, s.__length);
	p[s.__length] = 0;
	return p;
}

void *_cgo_ca7c2c1b0b09_Cfunc_CBytes(struct __go_open_array b) {
	char *p = malloc(b.__count);
	memmove(p, b.__values, b.__count);
	return p;
}

struct __go_string _cgo_ca7c2c1b0b09_Cfunc_GoString(char *p) {
	intgo len = (p != NULL) ? strlen(p) : 0;
	return __go_byte_array_to_string(p, len);
}

struct __go_string _cgo_ca7c2c1b0b09_Cfunc_GoStringN(char *p, int32_t n) {
	return __go_byte_array_to_string(p, n);
}

Slice _cgo_ca7c2c1b0b09_Cfunc_GoBytes(char *p, int32_t n) {
	struct __go_string s = { (const unsigned char *)p, n };
	return __go_string_to_byte_array(s);
}

extern void runtime_throw(const char *);
void *_cgo_ca7c2c1b0b09_Cfunc__CMalloc(size_t n) {
        void *p = malloc(n);
        if(p == NULL && n == 0)
                p = malloc(1);
        if(p == NULL)
                runtime_throw("runtime: C malloc failed");
        return p;
}

struct __go_type_descriptor;
typedef struct __go_empty_interface {
	const struct __go_type_descriptor *__type_descriptor;
	void *__object;
} Eface;

extern void runtimeCgoCheckPointer(Eface, Slice)
	__asm__("runtime.cgoCheckPointer")
	__attribute__((weak));

extern void localCgoCheckPointer(Eface, Slice)
	__asm__("main._cgoCheckPointer");

void localCgoCheckPointer(Eface ptr, Slice args) {
	if(runtimeCgoCheckPointer) {
		runtimeCgoCheckPointer(ptr, args);
	}
}

extern void runtimeCgoCheckResult(Eface)
	__asm__("runtime.cgoCheckResult")
	__attribute__((weak));

extern void localCgoCheckResult(Eface)
	__asm__("main._cgoCheckResult");

void localCgoCheckResult(Eface val) {
	if(runtimeCgoCheckResult) {
		runtimeCgoCheckResult(val);
	}
}
#line 1 "cgo-generated-wrappers"

