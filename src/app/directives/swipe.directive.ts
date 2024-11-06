import {
  Directive,
  ElementRef,
  EventEmitter,
  HostListener,
  Output
} from "@angular/core";

export type State = {
  t: number;
  startX: number;
  startY: number;
  x: number;
  y: number;
  xSpeed: number;
  ySpeed: number;
};

@Directive({
  selector: "[fsSwipe]",
  standalone: true,
})
export class SwipeDirective {
  @Output()
  onDragStart = new EventEmitter<State>();
  @Output()
  onDragEnd = new EventEmitter<State>();
  @Output()
  onDrag = new EventEmitter<State>();
  @Output()
  onSwipe = new EventEmitter<State>();

  private isHolding = false;
  private state: State;

  constructor(private host: ElementRef) {
    this.state = this.initialState(0, 0);
  }

  ngAfterViewInit(): void {
    const hostElement = this.host.nativeElement as HTMLElement;
    hostElement.style.cursor = "grab";
  }

  @HostListener("touchstart", ["$event"])
  @HostListener("mousedown", ["$event"])
  onMouseDown(event: MouseEvent) {
    event.preventDefault();
    this.isHolding = true;
    const hostElement = this.host.nativeElement as HTMLElement;
    hostElement.style.cursor = "grabbing";
    const { clientX, clientY } = event instanceof TouchEvent ? event.touches[0] : event;
    this.state = this.initialState(clientX, clientY);
    this.onDragStart.emit(this.state);
  }

  @HostListener("touchend", ["$event"])
  @HostListener("document:mouseup", ["$event"])
  onMouseUp(event: MouseEvent) {
    this.isHolding = false;
    const hostElement = this.host.nativeElement as HTMLElement;
    hostElement.style.cursor = "grab";
    this.onDragEnd.emit(this.state);
    event.stopImmediatePropagation();
  }
  
  @HostListener("touchmove", ["$event"])
  @HostListener("document:mousemove", ["$event"])
  onMouseOver(event: MouseEvent) {
    if (!this.isHolding) return;
    const { clientX, clientY } = event instanceof TouchEvent ? event.touches[0] : event;
    this.state = this.updatedState(this.state, clientX, clientY);
    this.onDrag.emit(this.state);
  }

  private initialState(x: number, y: number) {
    return {
      t: Date.now(),
      startX: x,
      startY: y,
      x,
      y,
      xSpeed: 0,
      ySpeed: 0,
    };
  }

  private updatedState(oldState: State, newX: number, newY: number) {
    const { t, x, y, startX, startY } = oldState;
    const newT = Date.now();
    const dt = newT - t;
    const dx = newX - x;
    const dy = newY - y;
    const xSpeed = dx / dt;
    const ySpeed = dy / dt;
    return {
      t: newT,
      x: newX,
      y: newY,
      xSpeed,
      ySpeed,
      startX,
      startY,
    };
  }
}
