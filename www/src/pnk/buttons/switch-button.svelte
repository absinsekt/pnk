<style src="./switch-button/switch-button.styl"></style>

<script>
  import { createEventDispatcher } from 'svelte';

  import Icon from 'pnk/icon/icon.svelte';

  import { spin } from './switch-button/switch-button';

  // dataGroup
  export let dataGroup;
  // label
  export let label = '';
  // size
  export let size = 'md';
  // icon
  export let icons = [];
  // icon
  export let index = 0;
  // block
  export let block = false;
  // disabled
  export let disabled = false;

  const dispatch = createEventDispatcher();
</script>

<button
  class="pnk-button"
  class:block={block}

  class:ico={label === ''}

  class:x1={size === 'sm'}
  class:x2={size === 'md'}
  class:x3={size === 'lg'}

  data-group={dataGroup}
  {disabled}

  on:click|preventDefault|stopPropagation={() => {
    index = index < icons.length-1 ? index+1 : 0;
    dispatch('change', index);
  }}
>

  <div class="icons">
  {#each icons as icon, idx}
    {#if idx === index}
    <div
      class="icon"
      transition:spin={{idx}}
    >
      <Icon {size} src={icon} />
    </div>
    {/if}
  {/each}
  </div>

  {label}
</button>
