# ✨ Magic Alias ✨

***This library is not recommended because its code is entirely AI-generated.***

> 🪄 A simple and powerful shell alias manager

## 🌟 What is Magic Alias?

Magic Alias (command: `ma`) is a friendly tool that helps you create and manage shell aliases. No more typing long commands over and over again!

## 🚀 Features

- 📝 **Create aliases easily**: Turn any long command into a short one
- 📋 **List all your aliases**: See all your shortcuts in one place
- 🗑️ **Remove aliases**: Delete shortcuts you don't need anymore
- 🖥️ **Interactive UI**: Nice menus make it easy to use
- 🔄 **Works with your shell**: Automatically sets up with your terminal

## 📦 Installation

```bash
# Install Magic Alias
go install github.com/akarachen/magic-alias/cmd/ma@latest

# Setup your shell automatically, with autocompletion included 🔋
ma init
```

## 🎮 How to Use

### Create a new alias

```bash
# Create with arguments
ma add g git

# Or use interactive mode
ma add
```

### List all your aliases

```bash
ma list
```

### Remove an alias

```bash
# Remove with arguments
ma remove g

# Or use interactive mode
ma rm
```

### Uninstall Magic Alias

```bash
# Uninstall
ma uninstall
```

## 🧰 Development

This project uses:
- [Cobra](https://github.com/spf13/cobra) for CLI commands
- [Charm](https://github.com/charmbracelet/huh) for interactive UI
- [Just](https://github.com/casey/just) for development tasks

To run tests:

```bash
just test
```

## 📄 License

See [LICENSE](LICENSE) for details.
