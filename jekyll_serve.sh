sudo docker run --rm --volume=$PWD:/srv/jekyll --network=host -it jekyll/jekyll:3.6 jekyll serve --watch --drafts
