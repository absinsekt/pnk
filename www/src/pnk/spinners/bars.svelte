<style src="./bars/bars.styl"></style>

<script>
  export let size = 'md';
</script>

<div class="pnk-spinner-bars"
  class:x1={size === 'sm'}
  class:x2={size === 'md'}
  class:x3={size === 'lg'}
>
  <div class="card1"></div>
  <div class="card2"></div>
  <div class="card3"></div>
  <div class="card4"></div>
  <div class="card5"></div>
</div>
