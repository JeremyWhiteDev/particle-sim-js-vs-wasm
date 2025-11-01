# Particle Simulator JS vs WASM

A basic web app to demonstrate differences between running basic loops/math calculations in WASM vs JS.

### Intro to Web Assembly

WASM, or WebAssembly, is a specification that was first shipped in browsers in 2017. It was a colllaborative effort between engineers from Mozilla, Google, Apple, Microsoft. The supposed primary goals of WASM were: running at near native speed, having a "portable" compile target for systems level coad (c, c++, Rust), and javascript/browser integration. 

Many web applications (Figma, Sketchup) rely on low level code (ex. c++) that is compiled to WASM to handle their graphics rendering. A common argument I found for leveraging WASM for these applications is becuase of the necessary performance tuning that comes with these engines. Another argument is that, while a rendering engine is possible to be built in any language (even Javascript) there is much better tooling/libraries that exist for these low level languages that have been used in native environments for many years/decades. While I admittadly don't know much about graphics engines, I can imagine that managing all the geometry and light/shading in an app like Sketchup can become incredibly demanding. I can also imagine how many of these concerns have been "solved problems" in native environments for some time, and how engineers might not want to reinvent the wheel in a scripting language like JS with all it's complexity (event loop, garbage collection) when there's proven libraries that abstract some of these problems away.

### My WASM experience

This is my very first project that I am targeting web assembly. I am writing this in go for starters and targetting WASM. Golang is a bit of a different animal because there is still a garbage collector running, and the targetted WASM will actually include the full go runtime. I like playing with go because I can pretty easily do what I want with the standard lib and I don't have to worry about integrating a testing framework or formatting rules.

### Why this project?

I was curious if I could write an application that I could actually see a difference in javascript vs wasm performance. I decided that, staying in the vein of rendering engines, I would create an application that is a "particle simulator" app, where we have many particles rendered to a html `<canvas>` element, and those particles move/collide/react to each other. I wanted to see if, even with naive implementations, there was a differnce in performance between JS and WASM given the same algorithms. I wasn't setting out to build a physics engine, just something that would do many computations per second.

### Hypothesis

For a simple particle simulator application (not a full physics engine) I expect to see _no difference_ between WASM and JS. I believe that the sorts of computations I will be doing, will be simple enough that the javascript engine that I use (V8) will be more than capable of running this logic.

### Results
...
