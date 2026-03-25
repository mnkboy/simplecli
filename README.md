# simplecli

CLI en Go con releases automáticos via GoReleaser y GitHub Actions.

## Instalación de GoReleaser

```bash
echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
sudo apt update
sudo apt install goreleaser
```

## Inicializar proyecto

```bash
goreleaser init
```

El archivo `.goreleaser.yaml` se genera automáticamente y configura cómo se compilarán los binarios para cada plataforma.

## Probar compilación local

```bash
goreleaser release --snapshot --clean
```

Los binarios se generan en la carpeta `dist/`.

## GitHub Actions

Archivo `.github/workflows/release.yml`:

```yaml
# This is an example .goreleaser.yml file with some sensible defaults.
# Make sure to check the documentation at https://goreleaser.com

# The lines below are called `modelines`. See `:help modeline`
# Feel free to remove those if you don't want/need to use them.
# yaml-language-server: $schema=https://goreleaser.com/static/schema.json
# vim: set ts=2 sw=2 tw=0 fo=cnqoj

version: 2

before:
  hooks:
    # You may remove this if you don't use go modules.
    - go mod tidy
    # you may remove this if you don't need go generate
    - go generate ./...

builds:
  - env:
      - CGO_ENABLED=0
    ldflags:
      - -s -w
      - -X main.version={{.Version}}
      - -X main.commit={{.Commit}}
      - -X main.date={{.Date}}
    goos:
      - linux
      - windows
      - darwin

archives:
  - formats: [tar.gz]
    # this name template makes the OS and Arch compatible with the results of `uname`.
    name_template: >-
      {{ .ProjectName }}_
      {{- title .Os }}_
      {{- if eq .Arch "amd64" }}x86_64
      {{- else if eq .Arch "386" }}i386
      {{- else }}{{ .Arch }}{{ end }}
      {{- if .Arm }}v{{ .Arm }}{{ end }}
    # use zip for windows archives
    format_overrides:
      - goos: windows
        formats: [zip]

changelog:
  sort: asc
  filters:
    exclude:
      - "^docs:"
      - "^test:"

release:
  footer: >-

    ---

    Released by [GoReleaser](https://github.com/goreleaser/goreleaser).

```

Cada vez que se suba un tag, GitHub Actions compilará automáticamente los binarios y los publicará.

## Makefile

```makefile
BINARY_NAME=simplecli

release:
	@echo "Creating release..."
	@read -p "Version: " version; \
	git tag $$version; \
	git push origin $$version
```

Para crear una nueva versión:

```bash
make release
```

Te pedirá la versión (ej: `v1.0.4`) y creará el tag automáticamente.

## Ver versión instalada

```bash
./simplecli --version
```

Salida esperada:
```
simplecli 1.0.3-beta
commit: 6fb6ec8bd70a77dc70171d3e2ebf6c6b56982ae8
built: 2026-03-09T04:49:28Z
```

## Comandos útiles

```bash
# Ver tags existentes
git tag

# Crear tag manualmente
git tag v1.0.5
git push origin v1.0.5

# Subir todos los tags
git push --tags
```

## Comandos utiles compilacion docker:

docker build -t simplecli:latest .
docker run --rm simplecli:latest

---

Para futuras actualizaciones, vuelve pronto para releer el archivo readme. Saludos !

---

## 🚀 Novedades del Release

Este release introduce un pipeline completo y automatizado de distribución para **simplecli**, facilitando su instalación y uso en múltiples plataformas.

---

### ✨ ¿Qué hay de nuevo?

* ⚙️ **Releases automatizados con GoReleaser**

  * Binarios multiplataforma (Linux, macOS, Windows)
  * Paquetes listos para usar (`.tar.gz`, `.zip`)
  * Generación automática de checksums para verificación

* 🔁 **Integración con CI/CD**

  * Los releases se generan automáticamente al crear un tag
  * Builds reproducibles y consistentes

* 🍺 **Soporte para Homebrew (Linux y macOS)**

  * Instalación con un solo comando:

    ```bash
    brew tap mnkboy/simplecli
    brew install simplecli
    ```
  * Actualizaciones sencillas con `brew upgrade`

---

### 🧪 Mejoras durante este release

Durante la implementación del sistema de distribución se realizaron varias mejoras importantes:

* Validación de accesibilidad pública de los binarios
* Verificación de descargas mediante herramientas CLI (`curl`)
* Corrección de inconsistencias entre tags y assets en los releases
* Ajustes en la fórmula de Homebrew para asegurar instalaciones confiables

---

### 📦 Instalación

#### Usando Homebrew

```bash id="lq7n0y"
brew tap mnkboy/simplecli
brew install simplecli
```

#### Instalación manual

```bash id="q2e0h9"
curl -L https://github.com/mnkboy/simplecli/releases/latest/download/simplecli_Linux_x86_64.tar.gz | tar xz
chmod +x simplecli
sudo mv simplecli /usr/local/bin/
```

---

### 🧠 ¿Por qué es importante?

Este release lleva a **simplecli** a un nivel listo para producción:

* 🔒 Builds reproducibles
* 🌍 Distribución pública
* ⚡ Instalación rápida y sencilla
* 🔄 Actualizaciones fáciles

---

### 🎯 Estado actual

**simplecli ya puede instalarse, actualizarse y distribuirse como una herramienta CLI profesional.**

---

### 💬 Feedback

Si encuentras algún problema o tienes sugerencias, no dudes en abrir un issue o contribuir al proyecto.
