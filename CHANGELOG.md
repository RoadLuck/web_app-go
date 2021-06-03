# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.1.0]
### Added
- Se añade la carpeta middleware para la configuración del CRSFToken.
- Se integra chi, como paquete externo para administrar las rutas.
- Se incorpora nosurf, el paquete que nos ayuda a manejar CRSDTokens.
- Se incorpora el paquete SCS para el manejo de sesiones.
- Se añadieron el template layout para evitar repeticiones en los templates.
- Se añadio el paquete config, para realizar las configuraciones generales de la app.
- Se añadio la carpeta models, en la cual se creo templatedata, para llevar la data desde la logica en los handlers al template. Ademas aqui tambien se alamacenaran modelos de DB.
### Changed
- Se cambio la forma en que se procesan los templates.


## [1.0.0] - 2021-07-27
### Added
- Se añadio el formato de carpetas cmd/web, pkg/handlers, pkg/render y templates.
- Se separa toda la logica de las carpetas del main a sus carpetas correspondientes.