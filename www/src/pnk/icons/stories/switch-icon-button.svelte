<script>
  import Button from 'pnk/buttons/button.svelte';
  import SwitchIcon from 'pnk/icons/switch-icon.svelte';

  import IcoClose from 'pnk/paths/close.svelte';
  import IcoMenu from 'pnk/paths/menu.svelte';

  let switcher = null;
</script>

<Button
  on:click={() => switcher.toggle() }>

  <SwitchIcon
    bind:this={switcher}
    icons={[IcoMenu, IcoClose]}
  />
</Button>
