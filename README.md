![New Version Planned](https://img.shields.io/badge/[NEW_VERSION_PLANNED]-007934?style=flat&logoColor=white)<BR>
**OCTOBER 2026**<br>
I'm planning a new version for gDroneMapper, better UI, better everything.  Locally run IA will be involved for pictures and videos analysis as well other goodies.  Yet to finish design before stepping into development.  Will update.  Peace.

----------

<img src="https://github.com/gershu-ar/gDroneMapper/blob/main/Screenshot_1.png?raw=true" height="300"> <img src="https://github.com/gershu-ar/gDroneMapper/blob/main/Screenshot_3.png?raw=true" width="500">

<a name="up"></a>

****************
## INDEX / ÍNDICE
### 🇺🇸 English<br>
- [Intro](#english)
- [What is gDroneMapper?](#gdrone)
- [Key features](#keyf)
- **[Download](#download)**
- [Support / Feedback](#support)
### 🇦🇷 Español
- [Introducción](#espanol)
- [¿Qué es gDroneMapper?](#gdrone_es)
- [Características claves](#keyf_es)
- **[Descargar](#download_es)**
- [Soporte / Comentarios](#support_es)
****************

<a name="english"></a>
# gDroneMapper - 🇺🇸 English

Current version: **1.0 Final**<br>
Published: **19/OCT/2025**

❔ Have you ever taken a drone picture and later asked yourself, **_"What the [phoque](https://www.youtube.com/watch?v=OMuHtTGqu6o) is that on the horizon?"_** — unsure of the camera's precise orientation and never really getting a clear answer?<br><br>
❓ Are you fed úp of manually plotting hundreds of pictures and videos onto a map to reconstruct a serious forensic analysis or to solve a problem only to discover you cannot export the information without **paying first**?<br><br>
⁉️ Is the task of mapping terrain and keeping track of hundreds of generated files making your work tedious and cumbersome while your cat demands his litter to be cleaned **[right meow](https://www.youtube.com/watch?v=5EkwBEL4MZw)**?

<a name="gdrone"></a>
## 🚁 gDroneMapper is here to save the day.
It processes DNG, JPG, and MP4 files in batch by folder, extracting rich telemetry data to generate a unique, clustered map. All the collected information is visualized under numbered pins, outputting a simple, highly compatible, portable, single-file HTML map.  It also generates in-map direct links to the absolute paths where the original files are located.

> [!TIP]
> Click here for a live example of a generated map: [DEMOSTRATION](https://gershu.ar/playground/gDroneMapper_20251019_183049.html)<br>

> [!NOTE] 
> Tested with files produced by a DJI Air 3S.  gDroneMapper should be compatible with all DJI generated pictures and videos.  Other brands too if standarized EXIF/XMP information is present.

[[Up]](#up)

<a name="keyf"></a>
## 📍 Key Features:

- Visual Coverage: For each image file, heading lines and cones of view (POVs) are created and graphically represented. This is based on mathematical calculations of the extracted data, providing a sense of visual coverage and the precise center of the image for reference. (This feature is available only for pictures, not videos, due to limitations in DJI's proprietary format. However, video pins and core telemetry are still presented.)

- Information Index: An information table is generated, indexing all processed files with useful telemetry data for quick reference and navigation.

- Visibility Control: You can individually or globally toggle the visibility of cones of view and/or pins, as well as hide the entire information table.

- Easy Pin Identification: 🔵 JPG, 🟢 DNG & 🟠 MP4

> [!IMPORTANT]
> Due to a lack of information in the files (technical limitation), **panoramic shot files store no telemetry to create a proper cone of vision**.  Apart from that, they are indexed and processed properly.  

- Map Styles: Three different map style layers are available.

- Data Export: A log file is generated in "copy and paste" CSV format, ready to be imported into Excel or a database.

- Open and Free: It is open source, modern, portable, fast, and completely free.

- Simple Distribution: The application consists of just three files: the backend, the frontend, and the configuration file.

- Technology Stack / Architecture: The backend (processor) is coded and compiled in Go (EXE compressed with UPX); ExifTool[^1] runtimes are embeded and deployed on first run.  The frontend (GUI) is coded in C# + .NET 9.0 + WinForms (uncompressed EXE, all required .NET 9.0 runtimes are self contained in the executable file, thus making a larger file).

[[Up]](#up)

<a name="download"></a>
 ## 💾 Compiled versions - Let's get started! (Windows 10/11 x64)
- **INSTALLER VERSION** (recommended)<br>
  (MSI) Download from [Releases](https://github.com/gershu-ar/gDroneMapper/releases) and run to install on your computer.<br>This will generate a desktop shortcut for easy app access.<br><br>
- **MANUAL / PORTABLE VERSION**<br>
  (ZIP) Download from [Releases](https://github.com/gershu-ar/gDroneMapper/releases), extract the files to a folder of your choice, and run "gDroneMapper.exe" to use.<br>You can also run "forrest.exe" by itself (it's the backend), it will read the configuration from "config.json" and process it accordingly.

> [!TIP]
> If you are unsure or you want to fool around with the code, you can compile from the source code yourself (the /SOURCE/ folder contains both the back and frontend).

> [!WARNING]
> If Windows's SmartScreen or any similar prompt appears, simply bypass it. This is a safe application and contains no viruses.


[[Up]](#up)

 <a name="support"></a>
## 🐛 **Known bugs / Support / Feedback**<br>

Frontend:<br>
* A few minor validations and GUI retouches required: working on version 1.1;
* Not known bugs.

Backend:<br>
* Needs some polishing in error handling;
* Not known bugs.

> [!IMPORTANT]
> For **support / report an Issue** or to **send feedback / requests** use the [Issues](https://github.com/gershu-ar/gDroneMapper/issues) section, I usually reply same day.
> Thank you.

Coded with ❤️ in 🇦🇷 Argentina<br>
www.gershu.ar

gDroneMapper's source codes and compiled Windows executables are provided "as is".<br>
[^1]: ExifTool (https://exiftool.org) - Kudos to Phil Harvey for this awesome tool <3

⬆️ [[Up]](#up)


*********
<a name="espanol"></a>
# gDroneMapper - 🇦🇷 Español

Versión actual: **1.0 Final**<br>
Publicada: **19/OCT/2025**

❔ ¿Alguna vez tomaste una foto con tu dron y luego te preguntaste, "**_¿¿Pero a qué mxxxda estaba apuntando en ese momento??_**" — sin saber la orientación precisa de la cámara y sin poder obtener una respuesta clara?<br><br>
❓ ¿Estás cansado de trazar manualmente cientos de fotos y videos en un mapa para reconstruir un análisis forense 🔍 serio o para resolver un problema sólo para descubrir que **tienes que pagar para poder exportar la información**?<br><br>
⁉️ ¿La tarea de mapear terrenos 🗺️ y llevar un registro de cientos de archivos generados está volviendo tu trabajo tedioso y engorroso mientras tu gato te mira demandando que le limpies su cagadero?

 <a name="gdrone_es"></a>
## 🚁gDroneMapper está aquí para salvar el día.
Procesa archivos DNG, JPG y MP4 por carpeta, extrayendo datos de telemetría detallados para generar un mapa único y agrupado. Toda la información recopilada se visualiza bajo pines numerados, con una salida en un archivo HTML simple, altamente compatible y portable.  También genera dentro del mapa vínculos directos a los archivos de acuerdo a la ruta absoluta.

> [!TIP]
> Click aquí para acceder un ejemplo real de un mapa generado: [DEMOSTRACIÓN](https://gershu.ar/playground/gDroneMapper_20251019_183049.html)<br>

> [!NOTE]
> Testeado con archivos generados por un DJI Air 3S.  gDroneMapper debería ser compatible con todos los modelos de DJI.  Otras marcas también en tanto y en cuanto respeten la nomenclatura en el perfil EXIF/XMP de los archivos.

[[Subir]](#up)

<a name="keyf_es"></a>
## 📍 Características claves:

- Cobertura Visual: Para cada archivo de imagen, se crean y representan gráficamente líneas de rumbo y conos de visión (POVs). Esto se basa en cálculos matemáticos de los datos extraídos, proporcionando una idea de la cobertura visual y el centro preciso de la imagen como referencia. (Esta función solo está disponible para fotografías, no para videos, debido a limitaciones en el formato propietario de DJI. Sin embargo, los pines de video y la telemetría básica se presentan de igual manera.)

- Índice de Información: Se genera una tabla de información 📊 que indexa todos los archivos procesados con datos de telemetría útiles para una referencia y navegación rápidas.

- Control de Visibilidad: Es posible activar o desactivar la visibilidad de los conos de visión y/o los pines de forma individual o global, así como ocultar la tabla de información completa.

- Identificación Sencilla de Pines: 🔵 JPG, 🟢 DNG & 🟠 MP4

> [!IMPORTANT]
> Debido a una falta de información propia de los archivos (limitación técnica), **las fotografías panorámicas no guardan información de telemetría adecuada para generar el cono de visión correspondiente**.  Fuera de eso, los archivos son indexados y procesados con normalidad.

- Estilos de Mapa: Tres capas diferentes de estilo de mapa disponibles.

- Exportación de Datos: Se genera un archivo de registro en formato CSV de "copiar y pegar", listo para ser importado en Excel o una base de datos. 💾

- Abierto y Gratuito: Es de código abierto, moderno, portable, rápido y completamente gratuito. 🆓

- Distribución Simple: La aplicación consta de solo tres archivos: el backend, el frontend y el archivo de configuración.

- Tecnología / Arquitectura: El backend (forrest.exe, procesador) fue codeado y compilado en Go (EXE comprimido con UPX); ExifTool[^2] está embebido y se despliega en la primera ejecución.  El frontend (gDroneMapper.exe, GUI) fue codeado en C# + .NET 9.0 + WinForms (EXE sin comprimir, todos los runtimes de .NET 9.0 necesarios están embebidos en el ejecutable, por ello el archivo tan extenso).

[[Subir]](#up)

 <a name="download_es"></a>
## Versión compilada - ¡Empecemos! (Windows 10/11 x64)
- **VERSIÓN CON INSTALADOR** (recomendada)<br>
  (MSI) Descarga desde [Versiones](https://github.com/gershu-ar/gDroneMapper/releases) y ejecutalo para instalar en tu computadora.<br>Esto generará un acceso directo en tu escritorio para fácil acceso.<br><br>
- **VERSIÓN MANUAL / PORTABLE**<br>
  (ZIP) Descarga desde [Versiones](https://github.com/gershu-ar/gDroneMapper/releases), extrae los archivos en la carpeta de tu elección, y ejecuta "gDroneMapper.exe" para utilizar<br>También podés ejecutar "forrest.exe" por su cuenta (es el backend): leerá la configuración en config.json y ejecutará acorde.

> [!TIP]
> Si no estás segurx o quieres jugar con el código, puedes compilar el código tú mismx (la carpeta /SOURCE/ contiene el código fuente tanto para el back como para el frontend).

> [!WARNING]
> Si aparece Windows SmartScreen o cualquier aviso similar, simplemente ignóralo. Es una aplicación segura y no contiene virus.

[[Subir]](#up)

 <a name="support_es"></a>
## 🐛 **Errores conocidos / Soporte / Comentarios**<br>

Frontend:<br>
* Se requieren algunas validaciones y retoques en la GUI: ya trabajando en la versión 1.1;<br>
* Sin errores conocidos.

Backend:<br>
* Se necesita pulir un poco el reporte de errores;<br>
* Sin errores conocidos.

> [!IMPORTANT]
> Para **soporte / reportar un problema** o **enviar Comentarios / pedidos** servite de abrir una [Incidencia](https://github.com/gershu-ar/gDroneMapper/issues) aquí en GitHub, habitualmente respondo en el día.<br>
> Gracias.

Codeado con ❤️ en 🇦🇷 Argentina<br>
www.gershu.ar

Los códigos fuente y ejecutables compilados de gDroneMapper para Windows se proporcionan "tal cual".<br>
[^2]: ExifTool (https://exiftool.org) - Gracias Phil Harvey por esta increíble herramienta <3

[[Subir]](#up)
