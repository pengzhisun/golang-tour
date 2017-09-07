@echo off

set $BASE_DIR=%~dp0

rem navigators
DOSKEY root=cd %BASE_DIR%

rem tour operations
DOSKEY newtour=pushd . $T cd %BASE_DIR% $T bash new_tour.sh $* $T popd
DOSKEY runtour=pushd . $T cd %BASE_DIR% $T bash run_tour.sh $* $T popd
DOSKEY codetour=code $1\$1.go
DOSKEY stagetour=git add $1

rem shortcuts
DOSKEY n2=notepad2 $1
DOSKEY nct=pushd . $T cd %BASE_DIR% $T bash new_tour.sh $* $T popd $T code $1\$1.go
DOSKEY cpt=git commit -m "$*" $T git push

rem devenv
DOSKEY upenv=pushd . $T cd %BASE_DIR% $T notepad2 devenv.bat $T popd
DOSKEY ldenv=pushd . $T cd %BASE_DIR% $T devenv.bat $T popd

rem test
DOSKEY test1=echo "test1"