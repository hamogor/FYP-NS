# Null Spire

### Camera:
Loop all rooms to find where the hallfways are. On the first tile in the hallway/exit from that room
mark that tile as a "camera_center_and_scale_to_room_when_moved_to"

### Races:
1. Ikthah - #42963f // Advanced Aliens
2. Cobalt Clan - #0c4475 + #3b94c0 // Mercs
3. (The) Null

### Animation mapping:
Create a set of animations that all entities *could* have. then name them the files
the type of entity it is, then an _ and its keyword: player_idle, nullGrunt_idle etc. . 

Shooting will be done using circles that have light like this gif
https://twitter.com/EarlySunGames/status/1232622999590907904

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

Try and frame the generator so that there are (typically) 3 encounters between objectives. 
* One war between any number of factions



First level:
Can also just kill him, even after taking the reward for the mission
One single unit of a faction is behind enemy lines for some reason (can choose to help on his mission). 
If he is killed, he drops a contract for the mission. (He should be placed second room from spawn)
If player chooses to help him then he is added as an ally, and they go on a mission to secure something 
in the level through a series of goals introducing the player to various aspects bla bla bla...
The mission should have a final encounter of a skirmish between the two factions that the NPC ally is NOT. 
e.g. if ally is faction a, then there is a skirmish between faction b and c. 

#### Ally system: 
Depending on the result of the first level:
1. Killing  the ally to get both faction b and c rewards (reward acts as faction influence to spend on skill tree)
2. Siding with the ally to to gain influence with their faction (encounter outside the room with boss of ally) and he offers you the reward of 
his faction (to spend on skill tree) (faction a)

When damage is done to an enemy and an ally they will simply respond with "hey watch it (friendly fire)"
but if damage is just done to an ally outright then all of them will become hostile.

Need better incentive to switch often (large variety in rewards.)



### Ui:
https://www.gamasutra.com/blogs/IanMartin/20141218/232772/Making_Your_PC_Game_Resolution_Agnostic__Solving_the_Resolution_Problem.php
X range 800 - 1920 px
Y range 600 - 1080 px

Everything needs to be based off the center of the screen.
Figure out the users native resolution, then figure out where the center
of the screen is.
All UI elements should fit inside a 800x600 space. So no matter the size of the screen, the interface
still works. All of the graphics should be relative to the centerX centerY location you get by
DesktopResX/2-1, DesktopResY/2-1

Then figure out where the edges are. The top left is going to be 0,0 in most cases
Bottom right is WWidth-1, WHeight-1

' Calculate sprite scale.
SpriteScale = Fix((Screen.Width / 640 + Screen.Height / 360) / 2)
If SpriteScale < 1 Then SpriteScale = 1
https://www.youtube.com/watch?v=U5skX_XzPWo&feature=youtu.be

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

