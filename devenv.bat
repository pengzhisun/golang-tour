@echo off

set $BASE_DIR=%~dp0

DOSKEY root=cd %BASE_DIR%

DOSKEY n2=notepad2 $1

DOSKEY newtour=pushd . $T cd %BASE_DIR% $T bash new_tour.sh $* $T popd
DOSKEY runtour=pushd . $T cd %BASE_DIR% $T bash run_tour.sh $* $T popd
DOSKEY codetour=code $1\$1.go
DOSKEY committour=git add $1 $T git commit -m "$2"