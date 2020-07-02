import {
  DateEquality,
  getStartOfMonth,
  isEqualDate,
} from 'pnk/core/date';

export type CalendarDay = {
  date: number;
  month: string;
  rawDate: Date;
  isToday: boolean;
  isSelected: boolean;
  isActiveMonth: boolean;
  isDisabled: boolean;
  isFirst: boolean;
  isLast: boolean;
}

export class CalendarConfig {
  constructor(
    public today = new Date(),
    public selectedDate: Date = null,
    public offsetDate: Date = null,
    public minDate: Date = null,
    public maxDate: Date = null,
    public isSundayFirst = false,
    public isWeekendDisabled = false,
  ) {}
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

export function getMonthCalendar(cfg = new CalendarConfig()): CalendarDay[] {
  const min = cfg.minDate.getTime();
  const max = cfg.maxDate.getTime();

  const result: CalendarDay[] = [];
  const position = getStartOfMonth(cfg.offsetDate);
  const offset = getDayOffset(position, cfg.isSundayFirst);

  position.setDate(-offset + 1);

  for (let i = 0; i < 42; i++) {
    const dt = new Date(position)
    const dtMills = dt.getTime();
    const dtDay = dt.getDay();
    const dtDate = dt.getDate();
    const dtMonth = dt.getMonth();

    result.push({
      date: dtDate,
      month: dtMonth.toString(), //TODO
      rawDate: dt,
      isToday: isEqualDate(dt, cfg.today),
      isSelected: cfg.selectedDate && isEqualDate(dt, cfg.selectedDate),
      isActiveMonth: isEqualDate(dt, cfg.offsetDate, DateEquality.Month),
      isDisabled: dtMills < min
        || dtMills > max
        || (cfg.isWeekendDisabled && (dtDay === 0 || dtDay === 6)),
      isFirst: isEqualDate(dt, cfg.minDate),
      isLast: isEqualDate(dt, cfg.maxDate),
    });

    position.setDate(position.getDate() + 1);
  }

  return result;
}
