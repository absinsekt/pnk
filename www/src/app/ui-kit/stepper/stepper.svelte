<style src="./stepper.styl"></style>

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

<div class="stepper">
  <div class="stepper-wrapper"
    class:off={value === 0}>
    <button
      on:click={flush}><i class="ico ico-close"></i></button>

    <input type="text" readonly
      bind:value
      on:click={decrement} />
  </div>

  <button
    on:click={increment}><i class="ico ico-plus"></i></button>
</div>
