import FormView from './views/form-view.svelte';

import Button from 'pnk/buttons/button.svelte';
import DatePicker from 'pnk/inputs/date.svelte';

import Suggest from 'pnk/lists/suggest.svelte';
import { dadataSuggest } from 'pnk/lists/suggest/providers/dadata';

import IcoCalendar from 'pnk/paths/calendar.svelte';

export default {
  title: 'Form',
};

export const simpleOrderForm = () => ({
  Component: FormView
});

export const suggest = () => ({
  Component: Suggest,
  props: {
    suggest: dadataSuggest,
    min: 5
  }
})

export const datePicker = () => ({
  Component: DatePicker,
  props: {
    value: new Date(),
  }
})

export const button = () => ({
  Component: Button,
  props: {
    size: 'sm',
    icon: IcoCalendar
  }
})
