import { Component, effect, HostListener, input } from "@angular/core";
import { Ad } from "../../models/responses";

@Component({
  selector: "fs-ad",
  imports: [],
  templateUrl: "./ad.component.html",
  styleUrl: "./ad.component.css"
})
export class AdComponent {
  ad = input.required<Ad>();
  imageIndex = 0;

  constructor() {
    effect(() => {
      this.imageIndex = 0;
      /* Preload images for better UX */
      this.ad().images.forEach(({ image_url }) => {
        const img = new Image();
        img.src = image_url;
      })
    });
  }

  @HostListener('click', ['$event'])
  onClick(event: MouseEvent) {
    const { clientX, clientY } = event;
    const targetRect = (event.target as HTMLElement)
      .getBoundingClientRect();
    const left = targetRect.x + targetRect.width / 2 - clientX > 0;
    setTimeout(() => {
      if (left) {
        this.imageIndex = this.imageIndex + this.ad().images.length - 1;
      } else {
        this.imageIndex = this.imageIndex + 1;
      }
      this.imageIndex %= this.ad().images.length;
      console.log(this.ad().id, this.ad().images[this.imageIndex].image_url);
    });
    event.stopPropagation();
  }
}
