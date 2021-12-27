package optimization

func OptimizeFrontBack(dictionary []string) []string {
    /*
    * Remove duplicate elements in the dictionary by making an array of the already found passwords and checking if the current password is already in it
    **/
    keys := make(map[string]bool);
    slice := []string{};
 
    for _, password := range dictionary {
        if _, has := keys[password]; !has {
            keys[password] = true;
            slice = append(slice, password);
        }
    }

    return slice;
}