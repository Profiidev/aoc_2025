{ pkgs, lib, ... }:

let
  libs = with pkgs; [
    lp_solve
  ];
in
{
  packages = libs;
  env.LD_LIBRARY_PATH = "${lib.makeLibraryPath libs}:$LD_LIBRARY_PATH";
  env.CGO_CFLAGS = "-I${pkgs.lp_solve}/include/lpsolve";
}
