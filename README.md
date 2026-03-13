# GUI Architecture Patterns with Fyne

This project demonstrates how to implement common GUI architecture patterns using the [Fyne GUI toolkit](https://fyne.io/) in Go.

## Overview

To illustrate the differences between these patterns, I have implemented a simple **To-Do application** three times, following these architectures:

- Model-View-Controller (MVC)
- Model-View-Presenter (MVP)
- Model-View-ViewModel (MVVM)

The repository is organized as a **Go workspace**, with each implementation contained in its own module. Within each module:

Because definitions for these patterns often vary, this project follows the interpretations and examples provided by **Arjan Codes** in the video: [Which Software Architecture Should You Use: MVC, MVP, or MVVM?](https://github.com/ArjanCodes/2022-gui).

Each implementation of the todo app includes the following functionality:

- **Task Visualization:** Display a list of current tasks.
- **Task Management:** Add new tasks and delete existing ones.
- **Persistence:** Task data is saved and loaded across sessions.

The persistence is implemented with a sqlite database in all variants.
