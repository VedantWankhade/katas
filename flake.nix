{
  description = "Python development environment";

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
            python313
            python313Packages.pytest
        ];

        shellHook = ''
          echo "🐍 Python development environment"
          echo "Happy Hacking 😎"
        '';
      };
    };
}