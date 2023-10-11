source_filename = "lib/builtin/print"

%type.string = type { i32, i32, i8* }

declare void @putchar(i8)

; print(string str) {
; 	  for int i = 0; i < str.size; i++ {
; 		  putchar(str.adr[i])
; 	  }
; }
define void @.print(%type.string* %str) {
entry:
	%i = alloca i32, align 8
	%0 = getelementptr inbounds %type.string, %type.string* %str, i32 0, i32 1
	%1 = getelementptr inbounds %type.string, %type.string* %str, i32 0, i32 2
	%2 = load i32, i32* %0, align 8
	%3 = load i8*, i8** %1, align 8
	store i32 0, i32* %i, align 8
	br label %for.cond

for.cond:
	%4 = load i32, i32* %i, align 8
	%str.len = getelementptr inbounds %type.string, %type.string* %str, i32 0, i32 1
	%5 = load i32, i32* %str.len, align 8
	%cmp = icmp slt i32 %4, %5
	br i1 %cmp, label %for.body, label %for.end

for.body:
	%str.adr = getelementptr inbounds %type.string, %type.string* %str, i32 0, i32 2
	%6 = load i8*, i8** %str.adr, align 8
	%7 = load i32, i32* %i, align 8
	%str.idx = getelementptr inbounds i8, i8* %6, i32 %7
	%8 = load i8, i8* %str.idx, align 1
	call void @putchar(i8 %8)
	br label %for.inc

for.inc:
	%9 = load i32, i32* %i, align 8
	%inc = add nsw i32 %9, 1
	store i32 %inc, i32* %i, align 8
	br label %for.cond

for.end:
	ret void
}

; println(string str) {
;     print(str)
;     putchar('\n')
; }
define void @.println(%type.string* %str) {
entry:
	call void @.print(%type.string* %str)
	call void @putchar(i8 10)
	ret void
}