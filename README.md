# FourStrongest Puzzle 

An **open-ended mathematical research project** exploring strategies to fit irregular 2D shapes optimally into a container.  
This project focuses on **packing, orientation, and placement optimization**, combining geometry, heuristics, and algorithmic efficiency.

---

## Problem Statement
Given a set of irregular 2D objects and a fixed container, determine an orientation and position for each object such that:
- All objects fit inside the container without overlap.
- The space utilization is maximized.

This problem is inspired by **real-world constraints** in logistics, manufacturing, and material science.

---

## Strategies Implemented
- **First-fit placement** → place each object in the first available valid position.  
- **Sweep algorithms** → horizontal and vertical sweeps across the container to find feasible placement.  
- **Shortest-path center placement** → position objects based on their geometric center and proximity to container boundaries.  
- **Constraint handling** → ensure objects fit fully inside the container with no overlap.  

---

## Tech Stack
- **Python** → prototyping algorithms and visualizations.  
- **Go** → optimized implementation for faster computation and scaling to larger instances.  

---

## Learnings
- Explored geometric algorithms and packing heuristics.  
- Compared Python prototyping vs. Go optimization in algorithm-heavy workloads.  
- Developed appreciation for computational geometry and optimization tradeoffs.  

---

## Setup
1. Clone the repo  
2. Run the Python prototype:  
   python solver.py
3. go run main.go

---

## Project Info
Type: Independent mathematical/algorithmic research
Languages: Python, Go
Status: Experimental project – open to extension (e.g. genetic algorithms, ML-based heuristics).



