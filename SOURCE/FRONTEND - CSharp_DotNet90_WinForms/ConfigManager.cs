using System;
using System.IO;
using System.Text.Json;
using System.Windows.Forms;

public static class ConfigManager
{
    private const string ConfigPath = "config.json";
    public static Config Current { get; private set; } = new();

    public static void Load()
    {
        if (!File.Exists(ConfigPath)) return;
        try
        {
            var json = File.ReadAllText(ConfigPath);
            var loaded = JsonSerializer.Deserialize<Config>(json);
            if (loaded != null) Current = loaded;
        }
        catch (Exception ex)
        {
            MessageBox.Show($"Failed to load config:\n{ex.Message}", "Error");
        }
    }

    public static void Save()
    {
        try
        {
            var json = JsonSerializer.Serialize(Current, new JsonSerializerOptions { WriteIndented = true });
            File.WriteAllText(ConfigPath, json);
        }
        catch (Exception ex)
        {
            MessageBox.Show($"Failed to save config:\n{ex.Message}", "Error");
        }
    }
}
