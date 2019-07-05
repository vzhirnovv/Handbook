# React lazy components and Suspense 

## ``React.lazy``

The ``React.lazy`` function lets you render a dynamic import as a regular component.

Using ``React.lazy`` will allow you to load the bundle containing the component when the component is
rendered.

``React.lazy()`` takes a function as an argument and returns a promise (dynamic import) which resolves
to a module with a default export containing the 'lazy' component.

## ``React.Suspense``

``React.Suspense`` provides a placeholder for the ``React.lazy`` component while it loads.

``React.Suspense`` accepts a ``fallback`` prop - a React element you want to render while waiting
for the component to load.
