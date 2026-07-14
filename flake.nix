{
  description = "aino_site";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-26.05";

    flake-utils.url = "github:numtide/flake-utils";

    treefmt-nix.url = "github:numtide/treefmt-nix";

    templ.url = "github:a-h/templ/v0.3.1020";
    templ.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      ...
    }@inputs:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };

        treefmtEval = inputs.treefmt-nix.lib.evalModule pkgs ./treefmt.nix;

        templ = inputs.templ.packages.${system}.templ;
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            templ
          ];

          shellHook = ''
            export PS1="(dev) $PS1"
          '';
        };

        formatter = treefmtEval.config.build.wrapper;
        checks = {
          formatting = treefmtEval.config.build.check self;
        };
      }
    );
}
