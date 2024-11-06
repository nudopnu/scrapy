import {
    AfterContentInit,
    Component,
    ElementRef,
    inject,
    Input,
    ViewChild,
} from "@angular/core";
import { ApiService } from "../../services/api.service";
import { Ad } from "../../models/responses";
import { CommonModule, JsonPipe } from "@angular/common";
import { State, SwipeDirective } from "../../directives/swipe.directive";
import { AdComponent } from "../../components/ad/ad.component";
import { MockAds } from "../../mock/ads.mock";
import { ActivatedRoute, Router } from "@angular/router";

@Component({
    selector: "fs-ads",
    standalone: true,
    imports: [JsonPipe, CommonModule, SwipeDirective, AdComponent],
    templateUrl: "./ads.component.html",
    styleUrl: "./ads.component.css",
})
export class AdsComponent implements AfterContentInit {
    @Input()
    agentId = "1";
    apiService = inject(ApiService);
    ads: Ad[] = [];

    @ViewChild("front")
    frontElementRef!: ElementRef;

    alpha = 0;
    frontTransform = "";
    backTransform = "translate3d(0, 0, -200px)";
    frontTransformOrigin = "";
    backTransformOrigin = "";
    transition = "";
    flipRot = 1;
    isSwipingOut = false;

    ngAfterContentInit(): void {
        console.log(this.agentId);
        this.apiService.getAdsByAgent(parseInt(this.agentId)).subscribe(
            (res) => {
                this.ads = res;
                setTimeout(() => {
                    this.frontTransform = "translate3d(0, 0, 0)";
                });
            },
        );
    }

    onDragStart(state: State) {
        const { startY } = state;
        const frontElement = this.frontElementRef.nativeElement as HTMLElement;
        const cy = frontElement.clientTop + frontElement.clientHeight / 2;
        const dy = startY - cy;
        if (dy < 0) {
            this.frontTransformOrigin = "bottom center";
            this.flipRot = 1;
        } else {
            this.frontTransformOrigin = "top center";
            this.flipRot = -1;
        }
        this.transition = "";
    }

    onDrag(state: State) {
        const { x, y, startX, startY } = state;
        const frontElement = this.frontElementRef.nativeElement as HTMLElement;
        const dx = x - startX;
        const dy = y - startY;
        const dxFrac = dx / frontElement.clientWidth;
        this.alpha = Math.min(1, Math.abs(dxFrac));
        const backTransZ = -200 * (1 - this.alpha);
        const frontRotZ = this.flipRot * dxFrac * 35;
        this.backTransform = `translate3d(0, 0, ${backTransZ}px)`;
        this.frontTransform = `translate3d(${dx * 0.7}px, ${
            dy * 1.3
        }px, 0) rotateZ(${frontRotZ}deg)`;
    }

    onDragEnd(state: State) {
        const { x, y, startX, startY, xSpeed } = state;
        const dx = x - startX;
        const dy = y - startY;
        const direction = dx > 0 ? 1 : -1;
        if (Math.abs(xSpeed) > 0.5) {
            this.isSwipingOut = true;
            setTimeout(() => {
                this.transition = "transform 0.4s ease";
                this.frontTransform = `translate3d(${
                    direction * document.body.clientWidth
                }px, ${dy * 1.3}px, 0) rotateZ(${
                    this.flipRot * direction * 45
                }deg)`;
                this.backTransform = `translate3d(0, 0, 0)`;
            });
        } else {
            setTimeout(() => {
                this.transition = "transform 0.2s ease";
                this.frontTransform = `translate3d(0, 0, 0) rotateZ(0)`;
                this.backTransform = `translate3d(0, 0, -200px)`;
            });
        }
    }

    onTransitionEnd() {
        this.transition = "";
        if (!this.isSwipingOut) return;
        this.backTransform = `translate3d(0, 0, -200px)`;
        this.frontTransform = `translate3d(0, 0, 0)`;
        this.ads = this.ads.slice(1);
        this.isSwipingOut = false;
    }

    mockObjects() {
        this.ads = MockAds;
    }
}
