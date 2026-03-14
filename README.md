# GUI Architecture Patterns with Fyne

This project demonstrates how to implement common GUI architecture patterns using the [Fyne GUI toolkit](https://fyne.io/) in Go.

## Overview

To illustrate the differences between these patterns, I have implemented a simple **To-Do application** three times, following these architectures:

- Model-View-Controller (MVC)
- Model-View-Presenter (MVP)
- Model-View-ViewModel (MVVM)

The repository is organized as a **Go workspace**, with each implementation contained in its own module.

Because definitions for these patterns often vary, this project follows the interpretations and examples provided by **Arjan Codes** in the video: [Which Software Architecture Should You Use: MVC, MVP, or MVVM?](https://github.com/ArjanCodes/2022-gui).

Each implementation of the todo app includes the following functionality:

- Display a list of current tasks.
- Add new tasks and delete existing ones.
- Task data is saved and loaded across sessions.
- An error dialog is shown when an error occurred.

The persistence is implemented with a sqlite database in all variants.

### GUI patterns

The following table summarizes the specific responsibilities assigned to each component across the three patterns:

| Pattern | Model (Data & Logic) | Logic Component (C/P/VM) | View (UI & Interaction) |
| --- | --- | --- | --- |
| **MVC** | Manages data and business logic; can notify the View of changes. | **Controller:** Handles user input and updates the Model accordingly. | Displays data from the Model; sends user input events to the Controller. |
| **MVP** | Manages data and business logic; remains independent of the View. | **Presenter:** Acts as a mediator; retrieves data from the Model and manually pushes updates to the View. | **Passive interface:** Forwards user events to the Presenter and is manually updated by it. |
| **MVVM** | Manages data and business logic; remains independent of the View. | **ViewModel:** Exposes state and commands; maintains synchronization with the View via **Data Binding**. | **Declarative UI:** Binds directly to the ViewModel; reacts automatically to data changes without manual intervention. |
