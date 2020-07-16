<style src="./switch-icon/switch-icon.styl"></style>

<script>
  import { createEventDispatcher } from 'svelte';

  import Icon from './icon.svelte';
  import { spin } from './switch-icon/switch-icon';

  // size
  export let size = 'md';
  // icons
  export let icons = [];

  const dispatch = createEventDispatcher();

  let index = 0;

  export function toggle() {
    index = index < icons.length-1 ? index+1 : 0;
    dispatch('change', index);
  }
</script>

<div class="pnk-switch-icon">
{#each icons as icon, idx}
  {#if idx === index}
  <div class="icon" transition:spin={{idx}}>
    <Icon {size} src={icon} />
  </div>
  {/if}
{/each}
</div>
