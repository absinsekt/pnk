<style src="./icon/icon.styl"></style>

<script>
  // rotated
  export let rotated = false;
  // size
  export let size = 'md';
  // src
  export let src;
</script>

<svg viewBox={`0 0 24 24`}
  class:rotated={rotated}

  class:x1={size === 'sm'}
  class:x2={size === 'md'}
  class:x3={size === 'lg'}>

  <svelte:component this={src} />
</svg>
