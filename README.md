# halogen
It is very REACTive

## build
```
export GOOS=js
export GOARCH=wasm
```

(in src folder)
```
go build -o yeet.wasm ./main.go  
```

## ideas
- primitives
- styling
  - default flex
  - rename
  - like figma
- table
  - column
    - visible
    - title
    - tooltip?
    - width
    - min width
    - fixed width
    - editable
    - sorting
    - alignment?
    - auto width
    - text wrap
- inputs

# style classes
- text
  - color
  - font
  - weight
  - line height
  - decorations
- layout
  - width (min, max)
  - height (min, max)
  - aspect ratio
  - direction
  - alignment
  - wrap
  - grid
  - columns
  - rows
  - span
  - postion
  - top, left, bottom, right
  - z-index/layers
  - overflow/scroll (color, width)
- box (maybe merge with layout)
  - visibility
  - color
  - radius
  - padding
  - margin
  - gap
  - opactiy
- border
  - color
  - weight
  - style
  - top, bottom, left, right
  - opactiy
- outline
  - color
  - weight
  - style
  - top, bottom, left, right
  - opactiy
- shadow
  - spread
  - opacity
  - color
- animation
  - delay
  - direction
  - duration
  - iterations
  - name
  - function
- interaction
  - selection
  - cursor
  - user select
  - pointer events
- border
- outline
- effects?
  - blur
- transform?
  - skew
  - scale
  - rotate

- hover
- active
- focus
- required
- disabled

# done
- width
- height
- padding
- margin

# naming
- aln = align?
- bkg = background?
- brd = border?
- flx = flex?
- gap = gap?
- grd = grid?
- hgt = height
- jst = justify?
- mar = margin
- ovr = overflow?
- pad = padding
- rot = rotation?
- txt = text/font?
- wdt = width
- wrp = wrap?
- shd = shadow? 

# groups
- bkg
  - color
- brd
  - color
  - width
  - style
- otl
  - color
  - width
  - style
  - margin
- flx
  - direction
  - wrap
- grd
  - columns
  - rows
- txt
  - color
  - weight
  - family
  - height
  - decoration
- lyt
  - span ? (maybe in grid)
  - overflow
  - position
  - left, right, top, bottom
  - aspect
- animations
  - transitions


- wth
  - max
- hgt
  - max
- pad
  - t
  - b
  - l
  - r
  - tb
  - lr
- mar
  - t
  - b
  - l
  - r
  - tb
  - lr
- jus
- aln
- gap
  - row
  - col
- dir
- wrp
- rad
  - tl
  - tr
  - bl
  - br



