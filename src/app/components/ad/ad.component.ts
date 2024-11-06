import { Component, Input, OnChanges } from "@angular/core";
import { Ad } from "../../models/responses";

@Component({
  selector: "fs-ad",
  standalone: true,
  imports: [],
  templateUrl: "./ad.component.html",
  styleUrl: "./ad.component.css",
})
export class AdComponent implements OnChanges {
  @Input()
  ad!: Ad;
  imageIndex = 0;
  
  ngOnChanges(): void {
    this.imageIndex = 0;
  }

  onClick(event: MouseEvent) {
    const { clientX, clientY } = event;
    const targetRect = (event.target as HTMLElement)
      .getBoundingClientRect();
    const left = targetRect.x + targetRect.width / 2 - clientX > 0;
    setTimeout(() => {
      if (left) {
        this.imageIndex = this.imageIndex + this.ad.images.length - 1;
      } else {
        this.imageIndex = this.imageIndex + 1;
      }
      this.imageIndex %= this.ad.images.length;
      console.log(this.imageIndex, this.ad.images.length);
    });
    event.stopPropagation();
  }
}
