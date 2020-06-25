<style src="./days.styl"></style>

<script>
  import { createEventDispatcher } from 'svelte';

  import Button from 'pnk/buttons/button';
  import IcoCalendarToday from 'pnk/paths/calendar-today.svelte';
  import IcoMenuLeft from 'pnk/paths/menu-left.svelte';
  import IcoMenuRight from 'pnk/paths/menu-right.svelte';

  import {
    WeekdaysShort,
    MonthsFull,
    getMonthCalendar,
    getOffsetDate,
  } from '../../date';

  const dispatch = createEventDispatcher();

  export let value;
  export let now;
  export let minDate;
  export let maxDate;
  export let offsetDate;
  export let dataGroup;

  $: calendar = getMonthCalendar(now, value, offsetDate, minDate, maxDate);

  $: isBackDisabled = calendar[0].isDisabled;
  $: isForwardDisabled = calendar[calendar.length - 1].isDisabled;

  function onChange(day) {
    if (day.isActiveMonth && !day.isDisabled) {
      value = new Date(day.rawDate);
      dispatch('change', value);
    }
  }
</script>

  <div class="pnk-date-header">
    <div class="pnk-date-date">

      <a class="pnk-dh-month" href="."
        on:click|preventDefault|stopPropagation={() => {}}>

        {MonthsFull[offsetDate.getMonth()]}</a>,

      <a class="pnk-dh-year" href="."
        on:click|preventDefault|stopPropagation={() => {}}>

        {offsetDate.getFullYear()}</a>
    </div>

    <div class="pnk-date-controls">
      <Button
        {dataGroup}
        size="sm"
        icon={IcoMenuLeft}
        disabled={isBackDisabled}
        on:click={() => offsetDate = getOffsetDate(offsetDate, 0, -1, 0)} />

      <Button
        {dataGroup}
        size="sm"
        icon={IcoCalendarToday}
        on:click={() => offsetDate = new Date(now)} />

      <Button
        {dataGroup}
        size="sm"
        disabled={isForwardDisabled}
        icon={IcoMenuRight}
        on:click={() => offsetDate = getOffsetDate(offsetDate, 0, 1, 0)} />
    </div>
  </div>

  <ul class="pnk-date-week">
    {#each WeekdaysShort as weekday}
    <li class="pnk-date-weekday">{weekday}</li>
    {/each}

    {#each calendar as day}
    <li class="pnk-date-day"
      class:today={day.isToday}
      class:selected={day.isSelected}
      class:month={day.isActiveMonth}
      class:disabled={day.isDisabled}

      on:click={() => onChange(day)}
    >

      {day.date}
    </li>
    {/each}
  </ul>
