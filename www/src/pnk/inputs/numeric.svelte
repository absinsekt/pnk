<style src="./numeric/numeric.styl"></style>

<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let value = 0;

  function increment() {
    dispatch('change', Number(value) + 1);
  }

  function decrement() {
    dispatch('change', Number(value) - 1);
  }

  function flush() {
    value = 0;
    dispatch('change', 0);
  }
</script>

<div class="numeric">
  <div class="numeric-wrapper"
    class:off={value === 0}>
    <button
      on:click={decrement}><i class="ico ico-minus"></i></button>

    <input type="text" readonly
      bind:value
      on:click={flush} />
  </div>

  <button
    on:click={increment}><i class="ico ico-plus"></i></button>
</div>
