<style src="./calendar.styl"></style>

<script>
  import Days from '../days/days.svelte';
  import Years from '../years/years.svelte';

  import {
    getOffsetDate,
  } from './calendar';

  export let now = new Date();
  export let selectedDate = new Date(now);
  export let minDate = getOffsetDate(now, 0, -2, 0);
  export let maxDate = getOffsetDate(now, 0, 2, 0);

  let mode = 0;
  let offsetDate = new Date(now);
</script>

<div class="pnk-calendar">
  {#if mode === 0}
    <Days
      {now}
      {selectedDate}
      {minDate}
      {maxDate}
      bind:mode
      bind:offsetDate
    />
  {:else if mode === 1}
    <Months />
  {:else}
    <Years />
  {/if}
</div>
