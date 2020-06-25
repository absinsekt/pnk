export const WeekdaysShort = ['пн', 'вт', 'ср', 'чт', 'пт', 'сб', 'вс'];
export const WeekdaysFull = ['понедельник', 'вторник', 'среда', 'четверг', 'пятница', 'суббота', 'воскресенье'];
export const MonthsFull = ['январь', 'февраль', 'март', 'апрель', 'май', 'июнь', 'июль', 'август', 'сентябрь', 'октябрь', 'ноябрь', 'декабрь'];

export type CalendarDay = {
  date: number;
  month: string;
  rawDate: Date;
  isToday: boolean;
  isSelected: boolean;
  isActiveMonth: boolean;
  isDisabled: boolean;
}

function getStartOfMonth(now: Date): Date {
  const result = new Date(now);

  result.setDate(1);

  return result;
}

function getDayOffset(date: Date, isSundayFirst: boolean = false): number {
  let result = isSundayFirst
    ? date.getDay()
    : date.getDay() - 1;

  if (result === -1) {
    result = 6
  }

  return result;
}

const enum DateEquality {
  Day = 0,
  Month = 1,
  Year = 2,
}

function isEqualDate(date1: Date, date2: Date, eqType:DateEquality = DateEquality.Day): boolean {
  if (date1.getFullYear() !== date2.getFullYear()) return false;
  if (eqType === DateEquality.Year) return true;

  if (date1.getMonth() !== date2.getMonth()) return false;
  if (eqType === DateEquality.Month) return true;

  return date1.getDate() === date2.getDate();
}

export function getOffsetDate(date: Date, days: number, months: number, years: number) {
  const result = new Date(date);

  result.setDate(result.getDate() + days);
  result.setMonth(result.getMonth() + months);
  result.setFullYear(result.getFullYear() + years);

  return result;
}

export function getMonthCalendar(today: Date, selectedDate: Date, offsetDate: Date, minDate: Date, maxDate: Date, isSundayFirst = false): CalendarDay[] {
  const min = minDate.getTime();
  const max = maxDate.getTime();

  const result: CalendarDay[] = [];
  const position = getStartOfMonth(offsetDate);
  const offset = getDayOffset(position, isSundayFirst);

  position.setDate(-offset + 1);

  for (let i = 0; i < 42; i++) {
    const dt = new Date(position)
    const dtMills = dt.getTime();
    const dtDate = dt.getDate();
    const dtMonth = dt.getMonth();

    result.push({
      date: dtDate,
      month: dtMonth.toString(), //TODO
      rawDate: dt,
      isToday: isEqualDate(dt, today),
      isSelected: selectedDate && isEqualDate(dt, selectedDate),
      isActiveMonth: isEqualDate(dt, offsetDate, DateEquality.Month),
      isDisabled: dtMills < min || dtMills > max
    });

    position.setDate(position.getDate() + 1);
  }

  return result;
}
