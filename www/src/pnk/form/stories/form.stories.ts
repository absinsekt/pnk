import FormView from './form-view.svelte';

export default {
  title: 'Form',
};

export const simpleOrderForm = () => ({
  Component: FormView
});

// import FormView from './views/form-view.svelte';

// import Button from 'pnk/buttons/button.svelte';
// import DatePicker from 'pnk/inputs/date-picker.svelte';

// import Suggest from 'pnk/lists/suggest.svelte';
// import { dadataSuggest } from 'pnk/lists/suggest/providers/dadata';

// import IcoCalendar from 'pnk/paths/calendar.svelte';

// import SwitchButton from 'pnk/buttons/switch-button.svelte';
// import IcoMenu from 'pnk/paths/menu.svelte';
// import IcoClose from 'pnk/paths/close.svelte';

// export default {
//   title: 'Form',
// };

// export const simpleOrderForm = () => ({
//   Component: FormView
// });

// export const suggest = () => ({
//   Component: Suggest,
//   props: {
//     suggest: dadataSuggest,
//     min: 5
//   }
// });

// export const datePicker = () => ({
//   Component: DatePicker,
//   props: {
//     value: new Date(),
//   }
// });

// export const button = () => ({
//   Component: Button,
//   props: {
//     size: 'sm',
//     icon: IcoCalendar
//   }
// });

// export const buttons = () => ({
//   Component: SwitchButton,
//   props: {
//     icons: [IcoMenu, IcoClose]
//   }
// });
