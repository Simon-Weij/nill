{
  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [ ];
      systems = [ "x86_64-linux" ];
      perSystem =
        {
          config,
          self',
          inputs',
          pkgs,
          system,
          ...
        }:
        {
          # temporary workaround.
          formatter = pkgs.writeShellScriptBin "formatter" ''
            if [[ $# = 0 ]]; then set -- .; fi
            exec "${pkgs.nixfmt-rfc-style}/bin/nixfmt" "$@"
          '';
          devShells = {
            default = pkgs.mkShell {
              buildInputs = [
                pkgs.gcc
                pkgs.go
                pkgs.git
              ];
            };
          };
        };

      flake = { };
    };
}
