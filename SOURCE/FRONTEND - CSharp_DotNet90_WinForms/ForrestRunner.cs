using System;
using System.Diagnostics;
using System.IO;

public static class ForrestRunner
{
    public static void Run(Action<string> outputHandler)
    {
        const string Executable = "forrest.exe";

        if (!File.Exists(Executable))
        {
            outputHandler("Executable not found: forrest.exe");
            return;
        }

        try
        {
            var process = new Process
            {
                StartInfo = new ProcessStartInfo
                {
                    FileName = Executable,
                    UseShellExecute = false,
                    RedirectStandardOutput = true,
                    RedirectStandardError = true,
                    CreateNoWindow = true
                }
            };

            process.OutputDataReceived += (s, e) => {
                if (!string.IsNullOrWhiteSpace(e.Data)) outputHandler(e.Data);
            };
            process.ErrorDataReceived += (s, e) => {
                if (!string.IsNullOrWhiteSpace(e.Data)) outputHandler(e.Data);
            };

            process.Start();
            process.BeginOutputReadLine();
            process.BeginErrorReadLine();
        }
        catch (Exception ex)
        {
            outputHandler($"Error running forrest.exe:\n{ex.Message}");
        }
    }
}
