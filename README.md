# Wirdos

WIP

## Technical Design

TODO: not sure on this, seems unnecessarily complicated...

### Scene

Beneath the conceptual model of a game, a scene would be the next largest structure, it is equivalent to a `Scene` in Unity or a `Room` in Game Maker Studio. My design is arranged against the idea of a film shoot or a theatre piece, while not perfect it helps outline responsibilities of the various components.

Everything in a scene is a arranged in layers to help the render pipeline understand what's to be rendered first and therefore in the background. A component in layer `-1` will be rendered later than a component in layer `0` so will appear in front should they occupy the same cartesian coordinates. If two components are in the same layer and overlap, their minimum y coordinate is used, the lower y coordinate is rendered afterwards so they remain in front.

#### Directors

While it makes little sense to have multiple directors in a theatre production, here it's used more to describe behaviours which direct other components in a scene but aren't visible. An example could be the `Player` which would be responsible for interpreting input and knowing which components to direct in response.

#### Actors

Visible components which are often highly dynamic in behaviour, while they're typically directed one or multiple `Directors`, they still own much of their physical occupation in the world. Interpreting collision with other `Actors` or `Set`, knowing how to react when colliding.

#### Props

Components in a scene which are interactable by characters. An example would be a remarkable object which the player could interact with and then make an observation for instance or potentially add to the character's inventory.

#### Blocks

Used to create the stage, can still be physically interactable; for instance, could be solid so a character can't move through it. Could be uninteractable, just decoration on the ground or in the foreground.

## Collision

All components in a `Scene` which are to interact with other components require colliders, these are tested every frame to check whether they're colliding with other colliders attached to other components.

A collider belongs to a `Body` which represents the physical attributes of a component within a `Scene`, bodies that are considered `static` so don't move or change ever could have colliders but should not check collision for themselves. They are to be checked against by `dynamic` bodies which can move within a `Scene`.

If a body is considered `solid`, it can't be moved through or move through other `solid` bodies. If a body is `solid` it is required to have a collider.
