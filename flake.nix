{
  description = "Exercism development environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux"; 
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
            exercism
            go_1_26
            golangci-lint
            python313
            python313Packages.pytest
        ];

        shellHook = ''
          echo "🏋️‍♂️ Exercism development environment"
          exercism configure -w `pwd`
          echo "Happy Hacking 😎"
        '';
      };
    };
}
