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

For the space thats not available to the player, (width is 3)
dig tunnels so the player can blast through, have loads of wires
and pipes and shit for assets. crawl space..


Tech modules are dotted around rooms that can move heavy equipment
around the room / do other stuff


### Ui:
mouse := r.Window.MousePosition()

mat.Project(mouse)

fmt.Print(mouse, "\n")

Any time you need to show a keybinding, give two options "OK", "EDIT". Edit brings up a menu that displays 
ONLY the keybinds that were displayed. This way, players can slowly edit the keybinds to their liking. 
Also dont display what key is being overwritten unless its a bool CUSTOM one. Its ok to edit default keys without
confirmation.

### RPG Elements:

### Font:
LECO_2014_Alt_Regular (in windows appdata)

### Enemies:
Enemies change tile propeties (set on fire/acid etc)

One faction is robotic and has a "wake-up" animation.

