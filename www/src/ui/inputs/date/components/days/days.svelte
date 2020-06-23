<style src="./days.styl"></style>

<script>
  import {
    WeekdaysShort,
    MonthsFull,
    getMonthCalendar,
    getOffsetDate,
  } from '../calendar/calendar';

  export let now;
  export let selectedDate;
  export let minDate;
  export let maxDate;
  export let mode;
  export let offsetDate;

  $: calendar = getMonthCalendar(now, selectedDate, offsetDate, minDate, maxDate);
</script>

<div
  class="pnk-date-days">

  <div class="pnk-date-header">
    <div class="pnk-date-date">
      <a class="pnk-dh-month" href="."
        on:click|preventDefault|stopPropagation={() => mode = 1}>

        {MonthsFull[offsetDate.getMonth()]}</a>,

      <a class="pnk-dh-year" href="."
        on:click|preventDefault|stopPropagation={() => mode = 2}>

        {offsetDate.getFullYear()}</a>
    </div>

    <div class="pnk-date-controls">
      <button on:click={() => offsetDate = getOffsetDate(offsetDate, 0, 0, -1)}>&lt;&lt;</button>
      <button on:click={() => offsetDate = getOffsetDate(offsetDate, 0, -1, 0)}>&lt;</button>
      <button on:click={() => offsetDate = new Date(now)}>TODAY</button>
      <button on:click={() => offsetDate = getOffsetDate(offsetDate, 0, 1, 0)}>&gt;</button>
      <button on:click={() => offsetDate = getOffsetDate(offsetDate, 0, 0, 1)}>&gt;&gt;</button>
    </div>
  </div>

  <ul class="pnk-date-week">
    {#each WeekdaysShort as weekday}
    <li class="pnk-date-day">
      <span class="pnk-date-wrapper pnk-date-wrapper-weekday">{weekday}</span>
    </li>
    {/each}

    {#each calendar as day}
    <li class="pnk-date-day"
      class:today={day.isToday}
      class:selected={day.isSelected}
      class:month={day.isActiveMonth}
      class:disabled={day.isDisabled}
    >

      <span class="pnk-date-wrapper">{day.date}</span>
    </li>
    {/each}
  </ul>
</div>
