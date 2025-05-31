from shapely.geometry import Polygon
from shapely.affinity import rotate, translate
import matplotlib.pyplot as plt

def create_polygon(points):
    return Polygon(points)

def rotate_polygon(polygon, angle):
    return rotate(polygon, angle, origin='centroid')

def move_polygon(polygon, x_offset, y_offset):
    return translate(polygon, xoff=x_offset, yoff=y_offset)


def greedy_place(shapes, container_width, container_height, rotation_angles=[0, 90, 180, 270]):
    placed_shapes = []

    for shape in shapes:
        placed = False
        for angle in rotation_angles:
            rotated = rotate_polygon(shape, angle)
            minx, miny, maxx, maxy = rotated.bounds
            shape_width = maxx - minx
            shape_height = maxy - miny

            # Sweep across container area
            for y in range(0, int(container_height - shape_height), 5):
                for x in range(0, int(container_width - shape_width), 5):
                    candidate = move_polygon(rotated, x - minx, y - miny)

                    # Check collision with all placed shapes
                    if all(not candidate.intersects(p) for p in placed_shapes):
                        placed_shapes.append(candidate)
                        placed = True
                        break
                if placed:
                    break
            if placed:
                break

        if not placed:
            print("Shape could not be placed.")
    
    return placed_shapes


def plot_shapes(shapes, container_width, container_height):
    fig, ax = plt.subplots()
    for shape in shapes:
        x, y = shape.exterior.xy
        ax.fill(x, y, alpha=0.5)

    ax.set_xlim(0, container_width)
    ax.set_ylim(0, container_height)
    ax.set_aspect('equal')
    plt.grid(True)
    plt.show()



# Define some sample shapes
shape1 = create_polygon([(0,0), (13,0), (16,4), (16,11), (11,11),(11,16),(0,16)])
shape2 = create_polygon([(0,0), (11,0), (11,11), (7,16),(0,11)])
#shape3 = create_polygon([(0,0), (7.2,0), (7.2,8.9), (0,8.9)])
#shape4 = create_polygon([(0,0), (9.1,0), (9.1,7.2), (0,7.2)])

#shapes = [shape1, shape2, shape3, shape4]
shapes = [shape1, shape2]
# Place shapes into a 100x100 container
container_w, container_h = 100, 100
placed = greedy_place(shapes, container_w, container_h)

# Visualize
plot_shapes(placed, container_w, container_h)






