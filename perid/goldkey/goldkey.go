/*
Library for Perid Language
*/

package goldkey

/*

string value cast string data

  param:
    string
  return:
    string
  example:
    "foobar" => foobar  
*/
func CastString(str string) string{
	return str[1: len(str) - 1]
}
