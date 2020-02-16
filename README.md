# Null Spire

### Races:
1. Ikthah - #42963f // Advanced Aliens
2. Cobalt Clan - #0c4475 + #3b94c0 // Mercs
3. (The) Null

### Animation mapping:
Create a set of animations that all entities *could* have. then name them the files
the type of entity it is, then an _ and its keyword: player_idle, nullGrunt_idle etc. . 

### Tile mapping:
Create a set of static tiles, and duplicates that act as states (blood splatter / burnt / damage)
Maybe create a new animation type that has an initial and final sprite.

### Level generation:
Need to be able to access each room to place assets

### Ui:
"do you not just invert the matrix to transform screen to world for a 2x3?"

Laplace expansion, turn 2x3 matrix into 3x3, then invert it, then multiply mouse
vector by the inverse transform matrix, then check if the result is inside image

a= 2.854, b = 0, c = 380 etc etc

### RPG Elements:

### Font:
LECO_2014_Alt_Regular (in windows appdata)

