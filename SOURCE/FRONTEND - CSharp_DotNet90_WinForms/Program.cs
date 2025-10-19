using System;
using System.Diagnostics;
using System.Drawing;
using System.IO;
using System.Text.RegularExpressions;
using System.Windows.Forms;

class DroneLauncher : Form
{
    private RichTextBox consoleBox = null!;
    private NotifyIcon? trayIcon;
    private TableLayoutPanel layout = null!;
    private LinkLabel statusLink = null!;

    [STAThread]
    static void Main()
    {
        Application.EnableVisualStyles();
        Application.SetCompatibleTextRenderingDefault(false);
        ConfigManager.Load();
        Application.Run(new DroneLauncher());
    }

    public DroneLauncher()
    {
        Text = "gDroneMapper";
        try
		{
			var asm = System.Reflection.Assembly.GetExecutingAssembly();
			using var stream = asm.GetManifestResourceStream("DroneLauncher.app.ico");
			Icon = stream != null ? new Icon(stream) : SystemIcons.Application;
		}
		catch { Icon = SystemIcons.Application; }

	
	
				
        Width = 480;
        Height = 700;
        StartPosition = FormStartPosition.CenterScreen;
        BackColor = Color.White;
        MaximizeBox = false;
        FormBorderStyle = FormBorderStyle.FixedSingle;
        Resize += OnResize;

        BuildLayout();
        SetupConsole();
		
		Shown += (s, e) =>
		{
			var lang = ConfigManager.Current.language ?? "en";
			AppendConsole(Translator.Translate("welcome_line1", lang) + "\n");
			AppendConsole(Translator.Translate("welcome_line2", lang) + "\n");
			AppendConsole(Translator.Translate("welcome_line3", lang) + "\n");
		};
		SetupExitButton();
        SetupStatusBar();
        SetupTrayIcon();
    }

    private void BuildLayout()
	{
		var config = ConfigManager.Current;
		var lang = config.language ?? "en";
		int row = 0;

		layout = new TableLayoutPanel {
			Dock = DockStyle.Top,
			AutoSize = true,
			AutoScroll = false,
			ColumnCount = 3
		};
		layout.ColumnStyles.Add(new ColumnStyle(SizeType.Percent, 35));
		layout.ColumnStyles.Add(new ColumnStyle(SizeType.Percent, 50));
		layout.ColumnStyles.Add(new ColumnStyle(SizeType.Percent, 15));
		Controls.Add(layout);

		// RUN BUTTON
		var runButton = new Button {
			Text = Translator.Translate("run", lang),
			BackColor = Color.ForestGreen,
			ForeColor = Color.White,
			Height = 40,
			Dock = DockStyle.Fill,
			Font = new Font("Segoe UI", 10, FontStyle.Bold)
		};
		runButton.Click += (s, e) => ForrestRunner.Run(AppendConsole);
		layout.Controls.Add(runButton, 0, row);
		layout.SetColumnSpan(runButton, 3);
		row++;

		// <SEPARACION>
		AddSeparator(row++);
		
		// INPUT DIRECTORY + ICON
		UIBuilder.AddLabeledTextBox(layout, "inputDirectory", config.inputDirectory ?? "", val => config.inputDirectory = val, lang, ContentAlignment.MiddleRight, row++);
		
		// CHECKBOX: READ FOLDER?
		UIBuilder.AddCheckbox(layout, "readSubfolders", config.recurse, val => config.recurse = val, lang, row++);
		
		// <LINEA>
		AddSeparator(row++);
		
		// OUTPUT DIRECTORY SELECTOR + ICON
		UIBuilder.AddLabeledTextBox(layout, "outputDirectory", config.outputDirectory ?? "", val => config.outputDirectory = val, lang, ContentAlignment.MiddleRight, row++);
		
		// <LINEA>
		AddSeparator(row++);
		
		// LOG DIRECTORY + ICON
		UIBuilder.AddLabeledTextBox(layout, "log_output_folder", config.log_output_folder ?? "", val => config.log_output_folder = val, lang, ContentAlignment.MiddleRight, row++);
		
		// CHECKBOX: ENABLE LOGGING?
		UIBuilder.AddCheckbox(layout, "enable_logging", config.enable_logging, val => config.enable_logging = val, lang, row++);
		
		// <LINEA>
		AddSeparator(row++);
		
		// DROPBOX: LANGUAGE SELECTOR
		UIBuilder.AddLanguageSelector(layout, config, row++);
		
		// INPUT DISTANCE + DROPDOWN UNIT
		UIBuilder.AddDistanceAndUnit(layout, config, lang, row++);
		
		// CHECKBOX: DRAW HEADING
		UIBuilder.AddCheckbox(layout, "drawHeading", config.drawHeading, val => config.drawHeading = val, lang, row++);
		
		// CHECKBOX: SCAN IMAGES
		UIBuilder.AddCheckbox(layout, "scan_images", config.scan_images, val => config.scan_images = val, lang, row++);
		
		// CHECKBOX: SCAN VIDEOS
		UIBuilder.AddCheckbox(layout, "scan_videos", config.scan_videos, val => config.scan_videos = val, lang, row++);
		
		// LEGEND
		var legend = new Label {
			Text = Translator.Translate("autosave", lang),
			Dock = DockStyle.Fill,
			TextAlign = ContentAlignment.MiddleCenter,
			Font = new Font("Segoe UI", 9, FontStyle.Italic)
		};
		layout.Controls.Add(legend, 0, row);
		layout.SetColumnSpan(legend, 3);
		row++;

		// <SEPARACION>
		AddSeparator(row++);
	}


    private void AddSeparator(int row)
	{
		var separator = new Label {
			BorderStyle = BorderStyle.Fixed3D,
			Height = 2,
			Dock = DockStyle.Fill
		};
		layout.Controls.Add(separator, 0, row);
		layout.SetColumnSpan(separator, 3);
	}



    private void SetupConsole()
    {
        consoleBox = new RichTextBox {
            Multiline = true,
            ReadOnly = true,
            WordWrap = false,
            ScrollBars = RichTextBoxScrollBars.ForcedBoth,
            Dock = DockStyle.Top,
            Height = 200,
            Font = new Font("Consolas", 9),
            TabStop = true
        };
        Controls.Add(consoleBox);
    }






    private void SetupExitButton()
    {
        var exitButton = new Button {
            Text = Translator.Translate("exit", ConfigManager.Current.language ?? "en"),
            BackColor = Color.DarkRed,
            ForeColor = Color.White,
            AutoSize = true,
            Height = 40,
            Dock = DockStyle.Bottom,
            Font = new Font("Segoe UI", 10, FontStyle.Bold)
        };
        exitButton.Click += (s, e) => Application.Exit();
        Controls.Add(exitButton);
    }

    private void SetupStatusBar()
    {
        var text = "gDroneMapper 1.0 – Coded by gershu";
        int linkStart = text.IndexOf("gershu");

        statusLink = new LinkLabel {
            Text = text,
            LinkArea = new LinkArea(linkStart, "gershu".Length),
            TextAlign = ContentAlignment.MiddleCenter,
            Dock = DockStyle.Bottom,
            Font = new Font("Segoe UI", 9),
            LinkColor = Color.Blue
        };
        statusLink.LinkClicked += (s, e) =>
        {
            try { Process.Start(new ProcessStartInfo { FileName = "https://gershu.ar", UseShellExecute = true }); }
            catch { MessageBox.Show("Could not open link."); }
        };
        Controls.Add(statusLink);
    }

    private void SetupTrayIcon()
    {
        trayIcon = new NotifyIcon();
        try
		{
			var asm = System.Reflection.Assembly.GetExecutingAssembly();
			using var stream = asm.GetManifestResourceStream("DroneLauncher.app.ico");
			trayIcon.Icon = stream != null ? new Icon(stream) : SystemIcons.Application;
		}
		catch { trayIcon.Icon = SystemIcons.Application; }


        trayIcon.Text = "gDroneMapper";
        trayIcon.Visible = false;

        var menu = new ContextMenuStrip();
        var exitItem = new ToolStripMenuItem("Exit");
        exitItem.Click += (s, e) => Application.Exit();
        menu.Items.Add(exitItem);
        trayIcon.ContextMenuStrip = menu;

        trayIcon.MouseClick += (s, e) =>
        {
            if (e.Button == MouseButtons.Left)
            {
                Show();
                WindowState = FormWindowState.Normal;
                BringToFront();
                trayIcon.Visible = false;
            }
        };
    }

    private void OnResize(object? sender, EventArgs e)
    {
        if (WindowState == FormWindowState.Minimized)
        {
            if (trayIcon?.Icon != null)
            {
                Hide();
                trayIcon.Visible = true;
            }
            else
            {
                WindowState = FormWindowState.Normal;
            }
        }
        else
        {
            if (trayIcon != null) trayIcon.Visible = false;
        }
    }

    private void AppendConsole(string line)
    {
        if (string.IsNullOrWhiteSpace(line)) return;

        consoleBox.Invoke((MethodInvoker)(() =>
        {
            var match = Regex.Match(line, @"Map generated at: (.+\\gDroneMapper_\d+_\d+\.html)");
            if (match.Success)
            {
                var path = match.Groups[1].Value;
                var filename = Path.GetFileName(path);
                var displayLine = $"Map generated at: {path} Process completed.";

                int start = displayLine.IndexOf(filename);
                consoleBox.AppendText(displayLine + Environment.NewLine);
                consoleBox.Focus();

                if (start >= 0)
                {
                    int globalStart = consoleBox.Text.LastIndexOf(filename);
                    consoleBox.Select(globalStart, filename.Length);
                    consoleBox.SelectionColor = Color.Blue;
                    consoleBox.SelectionFont = new Font(consoleBox.Font, FontStyle.Underline);
                    consoleBox.SelectionProtected = true;

                    consoleBox.MouseClick += (s, e) =>
                    {
                        var pos = consoleBox.GetCharIndexFromPosition(e.Location);
                        if (pos >= globalStart && pos <= globalStart + filename.Length)
                        {
                            try { Process.Start(new ProcessStartInfo { FileName = path, UseShellExecute = true }); }
                            catch { MessageBox.Show("Could not open file."); }
                        }
                    };
                }
            }
            else
            {
                consoleBox.AppendText(line.Contains("\n") ? line : line + Environment.NewLine);
                consoleBox.Focus();
            }

            consoleBox.SelectionStart = consoleBox.Text.Length;
            consoleBox.ScrollToCaret();
            consoleBox.Refresh();
        }));
    }
}
