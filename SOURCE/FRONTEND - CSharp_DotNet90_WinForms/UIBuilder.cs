using System;
using System.Drawing;
using System.Windows.Forms;

public static class UIBuilder
{
    public static void AddLabeledTextBox(
        TableLayoutPanel layout,
        string labelKey,
        string initialValue,
        Action<string> onChange,
        string lang,
        ContentAlignment alignment,
        int row)
    {
        var label = new Label {
            Text = Translator.Translate(labelKey, lang),
            TextAlign = alignment,
            Dock = DockStyle.Fill,
            Font = new Font("Segoe UI", 9, FontStyle.Bold)
        };
        layout.Controls.Add(label, 0, row);

        var textBox = new TextBox {
            Text = initialValue,
            Dock = DockStyle.Fill,
            Font = new Font("Segoe UI", 9)
        };
        textBox.TextChanged += (s, e) =>
        {
            onChange(textBox.Text);
            ConfigManager.Save();
        };
        layout.Controls.Add(textBox, 1, row);

        var browseButton = new Button {
            Text = "...",
            Width = 30,
            Dock = DockStyle.Fill,
            Font = new Font("Segoe UI", 9, FontStyle.Bold)
        };
        browseButton.Click += (s, e) =>
        {
            using var dialog = new FolderBrowserDialog();
            if (dialog.ShowDialog() == DialogResult.OK)
            {
                textBox.Text = dialog.SelectedPath;
                onChange(dialog.SelectedPath);
                ConfigManager.Save();
            }
        };
        layout.Controls.Add(browseButton, 2, row);
    }

    public static void AddDistanceAndUnit(TableLayoutPanel layout, Config config, string lang, int row)
	{
		var label = new Label {
			Text = Translator.Translate("distance_threshold", lang),
			TextAlign = ContentAlignment.MiddleRight,
			Dock = DockStyle.Fill,
			Font = new Font("Segoe UI", 9, FontStyle.Bold)
		};
		layout.Controls.Add(label, 0, row);

		var distanceBox = new TextBox {
			Text = config.distance.ToString("0.0"),
			Width = 60,
			Font = new Font("Segoe UI", 9),
			Margin = new Padding(0, 0, 5, 0)
		};
		distanceBox.TextChanged += (s, e) =>
		{
			if (double.TryParse(distanceBox.Text, out double val) && val >= 0.1 && val <= 2500)
			{
				config.distance = Math.Round(val, 1);
				ConfigManager.Save();
			}
		};

		var unitBox = new ComboBox {
			Width = 60,
			Font = new Font("Segoe UI", 9),
			DropDownStyle = ComboBoxStyle.DropDownList,
			Margin = new Padding(0)
		};
		unitBox.Items.AddRange(new[] { "m", "km" });
		unitBox.SelectedItem = config.unit ?? "m";
		unitBox.SelectedIndexChanged += (s, e) =>
		{
			config.unit = unitBox.SelectedItem?.ToString() ?? "m";
			ConfigManager.Save();
		};

		var panel = new Panel {
			Dock = DockStyle.Fill,
			AutoSize = true
		};
		panel.Controls.Add(distanceBox);
		panel.Controls.Add(unitBox);

		distanceBox.Location = new Point(0, 0);
		unitBox.Location = new Point(distanceBox.Width + 5, 0);

		layout.Controls.Add(panel, 1, row);
		layout.Controls.Add(new Label(), 2, row); // spacer
	}


    public static void AddLanguageSelector(TableLayoutPanel layout, Config config, int row)
	{
		var label = new Label {
			Text = Translator.Translate("language", config.language ?? "en"),
			TextAlign = ContentAlignment.MiddleRight,
			Dock = DockStyle.Fill,
			Font = new Font("Segoe UI", 9, FontStyle.Bold)
		};
		layout.Controls.Add(label, 0, row);

		var comboBox = new ComboBox {
			Width = 100,
			Font = new Font("Segoe UI", 9),
			DropDownStyle = ComboBoxStyle.DropDownList
		};
		comboBox.Items.AddRange(new[] { "English", "Español" });
		comboBox.SelectedItem = config.language == "sp" ? "Español" : "English";

		comboBox.SelectedIndexChanged += (s, e) =>
		{
			var selected = comboBox.SelectedItem?.ToString();
			var newLang = selected == "Español" ? "sp" : "en";
			if (newLang != config.language)
			{
				config.language = newLang;
				ConfigManager.Save();
				Application.Restart();
			}
		};

		layout.Controls.Add(comboBox, 1, row);
		layout.Controls.Add(new Label(), 2, row); // spacer
	}


    public static void AddCheckbox(
        TableLayoutPanel layout,
        string key,
        bool initialValue,
        Action<bool> onChange,
        string lang,
        int row,
        ContentAlignment align = ContentAlignment.MiddleRight)
    {
        var label = new Label {
            Text = Translator.Translate(key, lang),
            TextAlign = align,
            AutoSize = true,
            Dock = DockStyle.Fill,
            Font = new Font("Segoe UI", 9)
        };

        var box = new CheckBox {
            Checked = initialValue,
            Dock = DockStyle.Left
        };
        box.CheckedChanged += (s, e) =>
        {
            onChange(box.Checked);
            ConfigManager.Save();
        };

        layout.Controls.Add(label, 0, row);
        layout.Controls.Add(box, 1, row);
        layout.Controls.Add(new Label(), 2, row); // spacer
    }
}
