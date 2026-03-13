# Implementing GUI architecture patterns with Fyne

## Introduction

- Fyne is one of the most popular GUI frameworks for Go
- Fyne does not force developers

## Motivation

- Fyne does not require developers to adhere to any particular architecture pattern
- Still useful to apply an architecture pattern in some cases
  - Easier maintenance of larger applications
  - Division of work in large teams
  - Improved testability of the UI

## Approach

- We implement the same application for each architecture pattern

- The application is a simple todo app with the following features:

  - Show list of tasks
  - Add and delete tasks
  - Tasks are persistent

- Each components (e.g. model) is defined their own Go package

- Each application is implemented as it's own Go module

## MVC

## MVP

## MVVM

## Comparison
