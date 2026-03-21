# GUI Architecture Patterns with Fyne

This projects explores how to implement common GUI architecture patterns with the [Fyne GUI toolkit](https://fyne.io/).

## Approach

We have re-implemented the same GUI application for each of the following GUI patterns:

- Model-View (MV)
- Model-View-Controller (MVC)
- Model-View-Presenter (MVP)
- Model-View-ViewModel (MVVM)

Our example application is a To-DO app with this functionality:

- Display a list of current tasks.
- Add new tasks and delete existing ones.
- Task data is saved and loaded across sessions.
- An error dialog is shown when an error occurred.

## Code structure

The repository is organized as a **Go workspace**, with each implementation contained in its own module.

Each layer (e.g. model) is in a separate GO package. The main function is assembling the layers from the different packages and starting it.
