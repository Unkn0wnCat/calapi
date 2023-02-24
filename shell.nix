{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    nativeBuildInputs = [ pkgs.gcc pkgs.go ];

    shellHook =
      ''
        go get
        ./setupEnv.sh
        ./downloadLibraries.sh
        rm -r download
      '';
}