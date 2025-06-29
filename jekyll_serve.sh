sudo docker run --rm --volume=$PWD:/srv/jekyll -p 4000:4000 -it jekyll/jekyll:3.6 jekyll serve --watch --drafts
