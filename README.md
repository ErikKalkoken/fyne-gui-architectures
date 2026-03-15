# GUI Architecture Patterns with Fyne

This projects shows how to apply common GUI architecture patterns when implementing a GUI application with the with the [Fyne GUI toolkit](https://fyne.io/).

## Approach

For this project we have re-implemented the same application for each of the following GUI patterns:

- Model-View (MV)
- Model-View-Controller (MVC)
- Model-View-Presenter (MVP)
- Model-View-ViewModel (MVVM)

Our example GUI application is a To-DO app with this functionality:

- Display a list of current tasks.
- Add new tasks and delete existing ones.
- Task data is saved and loaded across sessions.
- An error dialog is shown when an error occurred.

## Code structure

The repository is organized as a **Go workspace**, with each implementation contained in its own module.

All files are in the main package. Each file represents a layer of the respective GUI pattern, e.g. model.go container the code for the model.
