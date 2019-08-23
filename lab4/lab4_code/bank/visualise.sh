#!/usr/bin/env bash

process_frame() {
    neato -Tpng ${1} -o ${1%.*}.png
    echo "Processed frame ${d}"
}

for d in out/*;
do
    process_frame ${d} &
done

wait

ffmpeg -framerate 3 -i out/%010d.png output.mp4