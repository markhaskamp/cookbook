
_pattern.matches_ requires a regex pattern that matches the entire string

    String s = "eddie would go";  
    Pattern p = Pattern.compile("die");  
    p.matches();  // => __false__  
  
    Pattern p2 = Pattern.compile(".*die.*");  
    p2.matches();  // => __true__  


Use a compination of _pattern.matcher_, _pattern.find_, _pattern.group_ for regex groups

    String s = "always read the plaque";
    Pattern p = Pattern.compile("a.");
    p.matcher();
    while (p.find()) {
      System.out.println(p.group());  // "ay" // then "ad" // then "aq"
    }

    


_group(0)_ is the entire matched string

    String s = "at Reds 4 Pirates 0";
    Pattern p = Pattern.compile("at (\\w+?)\\s");
    Matcher matcher = p.matcher(s);

    matcher.find();
    Assert.assertEquals("at Reds ", matcher.group(0));  // true


example of multiple groups

    String s = "at Reds 4 Pirates 0";
    Pattern p = Pattern.compile("at (\\w+?)\\s\\d+\\s(\\w+)");
    Matcher matcher = p.matcher(s);

    Assert.assertTrue(matcher.find());
    Assert.assertEquals("Reds", matcher.group(1));
    Assert.assertEquals("Pirates", matcher.group(2));

