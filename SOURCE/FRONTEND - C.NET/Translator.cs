using System.Collections.Generic;

public static class Translator
{
    private static readonly Dictionary<string, Dictionary<string, string>> Labels = new()
    {
        ["en"] = new()
        {
            ["run"] = "Run Forrest, Run!",
            ["exit"] = "EXIT",
            ["inputDirectory"] = "Input directory",
            ["readSubfolders"] = "Read subfolders?",
            ["outputDirectory"] = "Output directory",
            ["log_output_folder"] = "Log folder",
            ["enable_logging"] = "Enable logging?",
            ["distance"] = "Distance",
            ["unit"] = "Unit",
            ["distance_threshold"] = "Line drawing",
            ["language"] = "Language",
            ["drawHeading"] = "Draw heading",
            ["scan_images"] = "Scan images",
            ["scan_videos"] = "Scan videos",
            ["show_cones"] = "Show cones",
            ["autosave"] = "All options are saved automatically.",
            ["English"] = "English",
            ["Español"] = "Español",
            ["welcome_line1"] = "🛸 Welcome to gDroneMapper – ready to fly.\n",
            ["welcome_line2"] = "📂 Select your input folder and output folders\nand then make Forrest run!.\n",
            ["welcome_line3"] = "⚙️ All options are saved automatically.\n\n> Thank you for reporting bugs to soy@gershu.ar"
        },
        ["sp"] = new()
        {
            ["run"] = "Corre Forrest, corre!",
            ["exit"] = "SALIR",
            ["inputDirectory"] = "Carpeta de entrada",
            ["readSubfolders"] = "¿Leer subcarpetas?",
            ["outputDirectory"] = "Carpeta de salida",
            ["log_output_folder"] = "Carpeta de logs",
            ["enable_logging"] = "¿Activar logs?",
            ["distance"] = "Distancia",
            ["unit"] = "Unidad",
            ["distance_threshold"] = "Distancia de línea",
            ["language"] = "Idioma",
            ["drawHeading"] = "Dibujar dirección",
            ["scan_images"] = "Leer imágenes",
            ["scan_videos"] = "Leer videos",
            ["show_cones"] = "Mostrar conos",
            ["autosave"] = "Todas las opciones se guardan automáticamente.",
            ["English"] = "English",
            ["Español"] = "Español",
            ["welcome_line1"] = "🛸 Bienvenido a gDroneMapper – listo para despegar.\n",
            ["welcome_line2"] = "📂 Seleccioná tu carpeta de entrada y carpeta de salida,\n¡luego haz correr a Forrest!.\n",
            ["welcome_line3"] = "⚙️ Todas las opciones se guardan automáticamente.\n\n> Gracias por reportar bugs a soy@gershu.ar"
        }
    };

    public static string Translate(string key, string lang)
    {
        if (Labels.TryGetValue(lang, out var map) && map.TryGetValue(key, out var value))
            return value;
        return key;
    }
}
